package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"

	// "github.com/go-acme/lego/v4/challenge/http01"
	// "github.com/go-acme/lego/v4/challenge/tlsalpn01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/relex/aini"
	ag "github.com/sunshine69/automation-go/lib"
	u "github.com/sunshine69/golang-tools/utils"
)

// Inventory parser block
var (
	HostsPattern    string
	InventoryPath   string
	MatchedHostsMap map[string]*aini.Host
	HostList        []string
	Inventory       *aini.InventoryData
	Vars            map[string]any = make(map[string]any)
)

func LoadInventory(inventoryPath string) {
	println("Args Length: ", len(os.Args))
	HostsPattern = os.Args[1]
	if len(os.Args) > 2 {
		InventoryPath = os.Args[2]
	} else {
		InventoryPath = "inventory/hosts.ini"
	}
	Inventory = u.Must(aini.ParseFile(inventoryPath))
	inventoryDir := filepath.Dir(inventoryPath)
	u.CheckErr(Inventory.AddVars(inventoryDir), "AddVars")
	MatchedHostsMap = u.Must(Inventory.MatchHosts(HostsPattern))
	HostList = u.MapKeysToSlice(MatchedHostsMap)

	if len(os.Args) > 3 {
		// Loads command line vars
		for _, item := range os.Args[3:] {
			_tmp := strings.Split(item, "=")
			key, val := strings.TrimSpace(_tmp[0]), strings.TrimSpace(_tmp[1])
			Vars[key] = val
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
	maps.Copy(Vars, u.StringMapToAnyMap(host.Vars)) // Get command line opts
	Vars = u.Must(ag.FlattenAllVars(Vars))

	myUser := MyUser{
		Email:   Vars["account_email"].(string),
		KeyPath: Vars["user_key_path"].(string),
	}

	if u.MapLookup(Vars, "action", "").(string) == "create_user" {
		// Create a user. New accounts need an email and private key to start.
		privateKey := u.Must(ecdsa.GenerateKey(elliptic.P256(), rand.Reader))
		myUser.Key = privateKey
		myUser.SavePrivateKey()
	} else {
		myUser.LoadPrivateKey()
	}

	config := lego.NewConfig(&myUser)
	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	switch u.MapLookup(Vars, "env", "") {
	case "dev":
		config.CADirURL = Vars["ca_dir_url"].(string) //"http://192.168.99.100:4000/directory"
	case "uat":
		config.CADirURL = lego.LEDirectoryStaging
	case "prod":
		config.CADirURL = lego.LEDirectoryProduction
	}
	config.Certificate.KeyType = certcrypto.RSA2048
	// A client facilitates communication with the CA server.
	client := u.Must(lego.NewClient(config))

	// We specify an HTTP port of 5002 and an TLS port of 5001 on all interfaces
	// because we aren't running as root and can't bind a listener to port 80 and 443
	// (used later when we attempt to pass challenges). Keep in mind that you still
	// need to proxy challenge traffic to port 5002 and 5001.
	// u.CheckErr(client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "5002")),"SetHTTP01Provider")
	// u.CheckErr(client.Challenge.SetTLSALPN01Provider(tlsalpn01.NewProviderServer("", "5001")),"SetTLSALPN01Provider")
	u.CheckErr(client.Challenge.SetDNS01Provider(&MaraDNSProvider{Vars["maradns_config_file"].(string)}), "SetDNS01Provider")

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
	fmt.Printf("%#v\n", certificates)

	// ... all done.
}

func main() {
	if u.FileExistsV2(os.Args[2]) != nil {
		// Run this command to embed
		// go-bindata -pkg main -o plays/letsencrypt/bindata.go -nomemcopy inventory-letsencrypt/...
		println("Extracting default inventory dir")
		for _, as := range AssetNames() {
			fmt.Printf("Restore %s\n", as)
			RestoreAssets(".", as)
		}
	}
	LoadInventory(os.Args[2] + "/hosts.ini")
	playHost(MatchedHostsMap[HostList[0]])
}
