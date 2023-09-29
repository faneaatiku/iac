package iac

import "os"

type AppAuth struct {
	PublicKey  string
	PrivateKey string
}

func LoadAuthFromEnv(prefix string) *AppAuth {
	auth := AppAuth{}
	public, found := os.LookupEnv(prefix + "APP_PUBLIC_KEY")
	if !found {
		return &auth
	}
	auth.PublicKey = public

	private, found := os.LookupEnv(prefix + "APP_PRIVATE_KEY")
	if !found {
		return &auth
	}

	auth.PrivateKey = private

	return &auth
}
