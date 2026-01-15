package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"os"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	u "github.com/sunshine69/golang-tools/utils"
)

func GenerateX509Keypair() (*rsa.PrivateKey, *x509.CertificateRequest) {
	// Define your custom fields
	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            []string{"US"},
			Organization:       []string{"My Company"},
			OrganizationalUnit: []string{"IT Department"},
			Locality:           []string{"New York"},
			CommonName:         "example.com",
		},
		DNSNames: []string{"example.com", "www.example.com"},
	}

	// Generate a private key for the CSR
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Create the CSR in DER format
	csrDER, _ := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	println(u.JsonDump(csrDER, ""))
	// 4. Parse the bytes back into the *x509.CertificateRequest struct
	csr := u.Must(x509.ParseCertificateRequest(csrDER))

	os.WriteFile("server.key", u.Must(MarshalPKCS8PrivatePEM(privateKey)), 0o600)
	os.WriteFile("server.crt", u.Must(MarshalPKIXPublicKeyPEM(&privateKey.PublicKey)), 0o600)
	os.WriteFile("server.csr", MarshalCSRPEM(csr), 0o600)
	return privateKey, csr
}
func MarshalCSRPEM(csr *x509.CertificateRequest) []byte {
	// Wrap the raw DER bytes (csr.Raw) in a PEM block
	csrBlock := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csr.Raw,
	}

	// Encode to memory (returns []byte)
	return pem.EncodeToMemory(csrBlock)
}
func MarshalPKIXPublicKeyPEM(pubKey *rsa.PublicKey) ([]byte, error) {
	// 1. Convert to DER-encoded PKIX (SubjectPublicKeyInfo) format
	pubDER, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	// 2. Wrap in a PEM block
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubDER,
	}

	// 3. Encode to memory
	return pem.EncodeToMemory(pubBlock), nil
}
func MarshalPKCS8PrivatePEM(privKey *rsa.PrivateKey) ([]byte, error) {
	// 1. Convert to DER-encoded PKCS#8 format
	privDER, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	// 2. Wrap in a PEM block
	privBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privDER,
	}

	return pem.EncodeToMemory(privBlock), nil
}
func main() {
	privateKey, csr := GenerateX509Keypair()

	client := u.Must(lego.NewClient(lego.NewConfig(&MyUser{Email: "a@b"})))

	request := certificate.ObtainForCSRRequest{
		Bundle:     true,
		PrivateKey: privateKey,
		CSR:        csr,
	}
	// client here is a lego client, certificate coming from import lego
	certificate := u.Must(client.Certificate.ObtainForCSR(request))
	println(u.JsonDump(certificate, ""))
}

type MyUser struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
	// Extra data -
	KeyPath string
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
