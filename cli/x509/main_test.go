package main

import "testing"

func TestGenerateX509(t *testing.T) {
	GenerateX509Keypair()
	// Validate openssl can parse them
	// openssl req -noout -text -verify -in server.csr
	// openssl pkey -noout -pubin -text -in server.crt
	// openssl rsa -noout -text -in server.key

}
