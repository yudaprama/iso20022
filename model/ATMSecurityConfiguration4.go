package model

// Configuration of the digital signatures if the security module is able to perform digital signatures with an asymmetric key.
type ATMSecurityConfiguration4 struct {

	// Maximum number of certificates in a certificate path, the security module is able to manage.
	MaximumCertificates *Number `xml:"MaxCerts,omitempty"`

	// Maximum number of cosigners, the security module is able to manage in a digital signature.
	MaximumSignatures *Number `xml:"MaxSgntrs,omitempty"`

	// Digital signature algorithm the security module is able to manage.
	DigitalSignatureAlgorithm []*Algorithm14Code `xml:"DgtlSgntrAlgo,omitempty"`
}

func (a *ATMSecurityConfiguration4) SetMaximumCertificates(value string) {
	a.MaximumCertificates = (*Number)(&value)
}

func (a *ATMSecurityConfiguration4) SetMaximumSignatures(value string) {
	a.MaximumSignatures = (*Number)(&value)
}

func (a *ATMSecurityConfiguration4) AddDigitalSignatureAlgorithm(value string) {
	a.DigitalSignatureAlgorithm = append(a.DigitalSignatureAlgorithm, (*Algorithm14Code)(&value))
}
