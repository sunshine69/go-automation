package main

import (
	"crypto"
	"crypto/rsa"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	u "github.com/sunshine69/golang-tools/utils"
)

func main() {
	privateKey, _, csr := u.GenerateX509Keypair(&rsa.PrivateKey{}, map[string]any{})

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
