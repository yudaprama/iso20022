package model

// Configuration of the cryptographic keys.
type ATMSecurityConfiguration2 struct {

	// Maximum number of symmetric keys the security module is able to manage.
	MaximumSymmetricKey *Number `xml:"MaxSmmtrcKey,omitempty"`

	// Maximum number of asymmetric keys the security module is able to manage.
	MaximumAsymmetricKey *Number `xml:"MaxAsmmtrcKey,omitempty"`

	// Maximum RSA key length (in number of bytes), the security module is able to manage.
	MaximumRSAKeyLength *Number `xml:"MaxRSAKeyLngth,omitempty"`

	// Maximum RSA root key length (in number of bytes), the security module is able to manage, if different from the maximum RSA key length.
	MaximumRootKeyLength *Number `xml:"MaxRootKeyLngth,omitempty"`
}

func (a *ATMSecurityConfiguration2) SetMaximumSymmetricKey(value string) {
	a.MaximumSymmetricKey = (*Number)(&value)
}

func (a *ATMSecurityConfiguration2) SetMaximumAsymmetricKey(value string) {
	a.MaximumAsymmetricKey = (*Number)(&value)
}

func (a *ATMSecurityConfiguration2) SetMaximumRSAKeyLength(value string) {
	a.MaximumRSAKeyLength = (*Number)(&value)
}

func (a *ATMSecurityConfiguration2) SetMaximumRootKeyLength(value string) {
	a.MaximumRootKeyLength = (*Number)(&value)
}
