package model

// Value of the public component of a RSA key.
type PublicRSAKey1 struct {

	// Modulus of the RSA key.
	Modulus *Max5000Binary `xml:"Mdlus"`

	// Public exponent of the RSA key.
	Exponent *Max5000Binary `xml:"Expnt"`
}

func (p *PublicRSAKey1) SetModulus(value string) {
	p.Modulus = (*Max5000Binary)(&value)
}

func (p *PublicRSAKey1) SetExponent(value string) {
	p.Exponent = (*Max5000Binary)(&value)
}
