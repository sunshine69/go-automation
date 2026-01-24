package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/challenge/tlsalpn01"

	// "github.com/go-acme/lego/v4/challenge/http01"
	// "github.com/go-acme/lego/v4/challenge/tlsalpn01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/relex/aini"
	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
)

// Inventory parser block. This can be copied to a new play
var (
	HostsPattern    string
	InventoryDir    string
	MatchedHostsMap map[string]*aini.Host
	HostList        []string
	Inventory       *aini.InventoryData
	CommandlineVars map[string]any = make(map[string]any)
)

func init() {
	_HostsPattern := flag.String("H", "", "Host pattern (glob based)")
	// For this app as we use go-bindata to embed that dir thus it is hardcoded as the initial extraction will create that dir
	// however u can rename it and change it using -i option
	_InventoryDir := flag.String("i", "inventory-letsencrypt", "Inventory dir which contains hosts files (ini and yaml generator plugin supported) and inventory data (group_vars, host_vars etc)")

	var extraVars u.ArrayFlags
	flag.Var(&extraVars, "e", "Extra vars to pass to inventory data, like -e action=create_user -e var1=value1")
	debug := flag.Int("v", 0, "Verbose Debug level, default 0")

	flag.Parse()

	HostsPattern, InventoryDir = *_HostsPattern, *_InventoryDir

	if u.FileExistsV2(InventoryDir) != nil {
		// Run this command to embed
		// go-bindata -pkg main -o plays/letsencrypt/bindata.go -nomemcopy inventory-letsencrypt/...
		println("Extracting default inventory dir")
		for _, as := range AssetNames() {
			fmt.Printf("Restore %s\n", as)
			RestoreAssets(".", as)
		}
		println("[INFO] Looks like it is first time you run. The inventory template has been generated. You have to examine the values and change it as required. Inventory directory is inventory-letsencrypt")
		os.Exit(0)
	}

	Inventory, MatchedHostsMap, HostList, CommandlineVars = ag.LoadInventory(InventoryDir, HostsPattern, extraVars)

	if *debug > 0 {
		fmt.Fprintf(os.Stderr, "InventoryDir: %s - HostsPattern: %s - extraVars: %s\n", u.JsonDump(InventoryDir, ""), u.JsonDump(HostsPattern, ""), u.JsonDump(extraVars, ""))
		if *debug > 1 {
			fmt.Fprintf(os.Stderr, "Loaded InventoryData: %s\n", u.JsonDump(Inventory, ""))
		}
	}
}

// End inventory parser block

// Below is the problem domain that might use the inventory data. In this eg this is letsencrypt

// You'll need a user or account type that implements acme.User
type MyUser struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
	KeyPath      string
}

func (_u *MyUser) GetEmail() string {
	return _u.Email
}
func (_u MyUser) GetRegistration() *registration.Resource {
	return _u.Registration
}
func (_u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return _u.Key
}

func (_u *MyUser) LoadPrivateKey() {
	_u.Key = u.Must(u.LoadPrivateKeyFromPEM(_u.KeyPath))
}

func (_u *MyUser) SavePrivateKey() {
	// 2. Serialize the private key into PKCS#8 DER format bytes
	// PKCS#8 is the modern, recommended standard for storing private keys.
	keyBytes := u.Must(x509.MarshalPKCS8PrivateKey(_u.Key))
	// 3. Encode the DER bytes into PEM format
	pemBlock := &pem.Block{
		Type:  "PRIVATE KEY", // The standard type for a PKCS#8 key
		Bytes: keyBytes,
	}
	// 4. Create or open the output file
	file := u.Must(os.Create(_u.KeyPath))
	defer file.Close()
	// 5.
	u.CheckErr(pem.Encode(file, pemBlock), "Write the PEM encoded data to the file")
}

func playHost(host *aini.Host) {
	// Any func playXXX will take a host to do and u need to populate vars in the next three lines
	Vars := u.StringMapToAnyMap(host.Vars)
	maps.Copy(Vars, CommandlineVars) // Get command line opts override it if required
	Vars = u.Must(ag.FlattenAllVars(Vars))
	// End populate Vars. Vars can be used and pass around from now on

	// Check current expired or not - check_cert_domain is like domain:port
	if check_cert_domain := u.MapLookup(Vars, "check_cert_url", "").(string); check_cert_domain != "" {
		days := u.Must(strconv.Atoi(u.MapLookup(Vars, "days_to_expire", "10").(string)))
		if needUpdate, err := u.CheckCertExpiry(check_cert_domain, days); !needUpdate && err == nil {
			println("days to expire still greater then settings. Default is 10, add inventory var days_to_expire to set it")
			os.Exit(0)
		}
	}
	myUser := MyUser{
		Email:   Vars["account_email"].(string),
		KeyPath: Vars["user_key_path"].(string),
	}

	if u.MapLookup(Vars, "action", "").(string) == "create-user" {
		println("Create a user. New accounts need an email and private key to start.")
		privateKey := u.Must(ecdsa.GenerateKey(elliptic.P256(), rand.Reader))
		myUser.Key = privateKey
		myUser.SavePrivateKey()
	} else {
		myUser.LoadPrivateKey()
	}

	config := lego.NewConfig(&myUser)
	if _t := u.MapLookup(Vars, "ca_dir_url", ""); _t != "" {
		config.CADirURL = _t.(string)
	} else {
		switch u.MapLookup(Vars, "env", "") {
		case "dev":
			config.CADirURL = Vars["ca_dir_url"].(string) //"http://192.168.99.100:4000/directory"
		case "uat":
			config.CADirURL = lego.LEDirectoryStaging
		case "prod":
			config.CADirURL = lego.LEDirectoryProduction
		}
	}

	config.Certificate.KeyType = certcrypto.RSA2048
	client := u.Must(lego.NewClient(config))

	switch u.MapLookup(Vars, "challenge_provider", "http01").(string) {
	case "http01":
		u.CheckErr(client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", u.MapLookup(Vars, "http_port", "5002").(string))), "SetHTTP01Provider")
	case "tls01":
		u.CheckErr(client.Challenge.SetTLSALPN01Provider(tlsalpn01.NewProviderServer("", u.MapLookup(Vars, "https_port", "5001").(string))), "SetTLSALPN01Provider")
	case "dns01":
		u.CheckErr(client.Challenge.SetDNS01Provider(NewMaraDNSProvider(&MaraDNSProvider{Vars: Vars}), dns01.AddDNSTimeout(300*time.Second)), "SetDNS01Provider")
	}

	// New users will need to register
	reg := u.Must(client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true}))

	myUser.Registration = reg

	var privateKey crypto.PrivateKey = nil
	var SavePrivateKey bool = false
	if path := u.MapLookup(Vars, "private_key_path", "").(string); path != "" {
		privateKey_, err := u.LoadPrivateKeyFromPEM(path)
		if err != nil { // Can not load, leave nil to auto geenrate and save it
			SavePrivateKey = true
			privateKey = nil
		} else { // Use existing one
			privateKey = privateKey_
		}
	} else {
		privateKey = u.Must(ecdsa.GenerateKey(elliptic.P256(), rand.Reader))
		Vars["private_key_path"] = "server.key" // give it a value
		SavePrivateKey = true
	}

	request := certificate.ObtainRequest{
		Domains:    u.SliceMap(strings.Split(Vars["domain"].(string), ","), func(s string) *string { s1 := strings.TrimSpace(s); return &s1 }),
		Bundle:     true,
		PrivateKey: privateKey,
	}
	certificates := u.Must(client.Certificate.Obtain(request))

	if SavePrivateKey {
		u.CheckErr(os.WriteFile(u.MapLookup(Vars, "private_key_path", "server.key").(string), certificates.PrivateKey, 0o600), "Save Private Key")
	}
	u.CheckErr(os.WriteFile(u.MapLookup(Vars, "public_key_path", "server.crt").(string), certificates.Certificate, 0o600), "Save Cert")
	// Each certificate comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL. SAVE THESE TO DISK.
	// fmt.Printf("%#v\n", certificates)
	if postCmd := u.MapLookup(Vars, "post_command", "").(string); postCmd != "" {
		o := u.Must(u.RunSystemCommandV2(postCmd, true))
		println(o)
	}
	// ... all done.
}

func main() { // This only play the first host. If multiple hosts resolved from the host pattern we can spawn go routine per hosts
	playHost(MatchedHostsMap[HostList[0]])
}
