package main

// IIS Certificate Update CLI (Pure Go, No PowerShell)
// ---------------------------------------------------
// Features:
// - Import PFX certificate (with private key) into LocalMachine\MY store
// - Extract correct certificate thumbprint from Windows cert store
// - Replace IIS HTTPS binding (SNI hostname binding)
// - No PowerShell, no Ansible, no external runtime dependencies
//
// REQUIREMENTS:
// - Run as Administrator
// - Windows Server with IIS
// - Go 1.21+ recommended
//
// Usage:
//   iis-cert-update.exe -host example.com -pfx cert.pfx -password secret

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	crypt32 = windows.NewLazySystemDLL("crypt32.dll")
	httpapi = windows.NewLazySystemDLL("httpapi.dll")

	procPFXImportCertStore                = crypt32.NewProc("PFXImportCertStore")
	procCertEnumCertificatesInStore       = crypt32.NewProc("CertEnumCertificatesInStore")
	procCertAddCertificateContextToStore  = crypt32.NewProc("CertAddCertificateContextToStore")
	procCertOpenStore                     = crypt32.NewProc("CertOpenStore")
	procCertGetCertificateContextProperty = crypt32.NewProc("CertGetCertificateContextProperty")

	httpSetServiceConfiguration    = httpapi.NewProc("HttpSetServiceConfiguration")
	httpDeleteServiceConfiguration = httpapi.NewProc("HttpDeleteServiceConfiguration")
)

const (
	CERT_STORE_PROV_SYSTEM          = 10
	CERT_SYSTEM_STORE_LOCAL_MACHINE = 0x20000

	CRYPT_MACHINE_KEYSET       = 0x20
	PKCS12_ALLOW_OVERWRITE_KEY = 0x4000

	CERT_HASH_PROP_ID = 3

	HttpServiceConfigSslSniCertInfo = 2
)

func main() {
	var (
		host     = flag.String("host", "", "IIS hostname binding (example.com)")
		pfxFile  = flag.String("pfx", "", "PFX certificate file")
		password = flag.String("password", "", "PFX password")
		port     = flag.Int("port", 443, "HTTPS port")
	)

	flag.Parse()

	if *host == "" || *pfxFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	thumb, err := importPFXToMachineStore(*pfxFile, *password)
	if err != nil {
		log.Fatal("Import failed:", err)
	}

	fmt.Println("Imported certificate thumbprint:", thumb)

	if err := updateIISSslBinding(*host, *port, thumb); err != nil {
		log.Fatal("Binding update failed:", err)
	}

	fmt.Println("IIS certificate binding updated successfully.")
}

// ------------------------------------------------------------------
// Import PFX into LocalMachine\MY and extract correct thumbprint
// ------------------------------------------------------------------
func importPFXToMachineStore(pfxPath, password string) (string, error) {
	pfxData, err := os.ReadFile(pfxPath)
	if err != nil {
		return "", err
	}

	blob := windows.CertBlob{
		Size: uint32(len(pfxData)),
		Data: &pfxData[0],
	}

	pw, _ := syscall.UTF16PtrFromString(password)

	store, _, err := procPFXImportCertStore.Call(
		uintptr(unsafe.Pointer(&blob)),
		uintptr(unsafe.Pointer(pw)),
		uintptr(CRYPT_MACHINE_KEYSET|PKCS12_ALLOW_OVERWRITE_KEY),
	)

	if store == 0 {
		return "", err
	}

	// Open LocalMachine\MY store
	storeName, _ := syscall.UTF16PtrFromString("MY")

	machineStore, _, err := procCertOpenStore.Call(
		CERT_STORE_PROV_SYSTEM,
		0,
		0,
		CERT_SYSTEM_STORE_LOCAL_MACHINE,
		uintptr(unsafe.Pointer(storeName)),
	)

	if machineStore == 0 {
		return "", err
	}

	var ctx uintptr

	for {
		ctx, _, _ = procCertEnumCertificatesInStore.Call(store, ctx)
		if ctx == 0 {
			break
		}

		procCertAddCertificateContextToStore.Call(
			machineStore,
			ctx,
			1,
			0,
		)

		// Extract SHA1 thumbprint
		var hash [20]byte
		var size uint32 = 20

		ret, _, _ := procCertGetCertificateContextProperty.Call(
			ctx,
			CERT_HASH_PROP_ID,
			uintptr(unsafe.Pointer(&hash[0])),
			uintptr(unsafe.Pointer(&size)),
		)

		if ret != 0 {
			return hex.EncodeToString(hash[:]), nil
		}
	}

	return "", fmt.Errorf("no certificate found in PFX")
}

// ------------------------------------------------------------------
// Update IIS HTTPS Binding via HTTP.sys SNI binding
// ------------------------------------------------------------------
func updateIISSslBinding(host string, port int, thumb string) error {
	hash, err := hex.DecodeString(thumb)
	if err != nil {
		return err
	}

	hostPort := fmt.Sprintf("%s:%d", host, port)
	hostPtr, _ := syscall.UTF16PtrFromString(hostPort)

	type HTTP_SERVICE_CONFIG_SSL_SNI_SET struct {
		KeyDesc struct {
			Host *uint16
		}
		ParamDesc struct {
			SslHash *byte
			AppId   windows.GUID
		}
	}

	cfg := HTTP_SERVICE_CONFIG_SSL_SNI_SET{}
	cfg.KeyDesc.Host = hostPtr
	cfg.ParamDesc.SslHash = &hash[0]
	cfg.ParamDesc.AppId = windows.GUID{
		Data1: 0x4dc3e181,
		Data2: 0xe14b,
		Data3: 0x4a21,
		Data4: [8]byte{0xb0, 0x22, 0x59, 0xfc, 0x66, 0x9b, 0x09, 0x14},
	}

	httpDeleteServiceConfiguration.Call(
		0,
		HttpServiceConfigSslSniCertInfo,
		uintptr(unsafe.Pointer(&cfg)),
		uintptr(unsafe.Sizeof(cfg)),
		0,
	)

	ret, _, err := httpSetServiceConfiguration.Call(
		0,
		HttpServiceConfigSslSniCertInfo,
		uintptr(unsafe.Pointer(&cfg)),
		uintptr(unsafe.Sizeof(cfg)),
		0,
	)

	if ret != 0 {
		return err
	}

	return nil
}
