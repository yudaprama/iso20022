package model

// Value of the public component of a RSA key.
type PublicRSAKey2 struct {

	// Asymmetric cryptographic algorithm.
	Algorithm *Algorithm7Code `xml:"Algo,omitempty"`

	// Public key value.
	PublicKeyValue *PublicRSAKey1 `xml:"PblcKeyVal"`
}

func (p *PublicRSAKey2) SetAlgorithm(value string) {
	p.Algorithm = (*Algorithm7Code)(&value)
}

func (p *PublicRSAKey2) AddPublicKeyValue() *PublicRSAKey1 {
	p.PublicKeyValue = new(PublicRSAKey1)
	return p.PublicKeyValue
}
