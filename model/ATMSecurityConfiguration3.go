package model

// Configuration of the encryption or digital envelope, if the security module is able to perform encryption.
type ATMSecurityConfiguration3 struct {

	// True if the security module is able to perform encryption with an asymmetric key.
	AsymmetricEncryption *TrueFalseIndicator `xml:"AsmmtrcNcrptn,omitempty"`

	// True if the security module is able to identify an asymmetric key with certificate issuer X.500 name and certificate serial number. False if a proprietary asymmetric key identifier is required.
	AsymmetricKeyStandardIdentification *TrueFalseIndicator `xml:"AsmmtrcKeyStdId,omitempty"`

	// Asymmetric encryption algorithm the security module is able to manage.
	AsymmetricEncryptionAlgorithm []*Algorithm7Code `xml:"AsmmtrcNcrptnAlgo,omitempty"`

	// True if the security module is able to manage a symmetric transport session key to protect cryptographic keys and data. False if only a previously exchanged symmetric key must be used; a proprietary symmetric key identifier is then used.
	SymmetricTransportKey *TrueFalseIndicator `xml:"SmmtrcTrnsprtKey,omitempty"`

	// Symmetric transport session key algorithm the security module is able to manage.
	SymmetricTransportKeyAlgorithm []*Algorithm13Code `xml:"SmmtrcTrnsprtKeyAlgo,omitempty"`

	// Symmetric encryption algorithm the security module is able to manage.
	SymmetricEncryptionAlgorithm []*Algorithm15Code `xml:"SmmtrcNcrptnAlgo,omitempty"`

	// Format of data before encryption, if the format is not plaintext or implicit.
	EncryptionFormat []*EncryptionFormat1Code `xml:"NcrptnFrmt,omitempty"`
}

func (a *ATMSecurityConfiguration3) SetAsymmetricEncryption(value string) {
	a.AsymmetricEncryption = (*TrueFalseIndicator)(&value)
}

func (a *ATMSecurityConfiguration3) SetAsymmetricKeyStandardIdentification(value string) {
	a.AsymmetricKeyStandardIdentification = (*TrueFalseIndicator)(&value)
}

func (a *ATMSecurityConfiguration3) AddAsymmetricEncryptionAlgorithm(value string) {
	a.AsymmetricEncryptionAlgorithm = append(a.AsymmetricEncryptionAlgorithm, (*Algorithm7Code)(&value))
}

func (a *ATMSecurityConfiguration3) SetSymmetricTransportKey(value string) {
	a.SymmetricTransportKey = (*TrueFalseIndicator)(&value)
}

func (a *ATMSecurityConfiguration3) AddSymmetricTransportKeyAlgorithm(value string) {
	a.SymmetricTransportKeyAlgorithm = append(a.SymmetricTransportKeyAlgorithm, (*Algorithm13Code)(&value))
}

func (a *ATMSecurityConfiguration3) AddSymmetricEncryptionAlgorithm(value string) {
	a.SymmetricEncryptionAlgorithm = append(a.SymmetricEncryptionAlgorithm, (*Algorithm15Code)(&value))
}

func (a *ATMSecurityConfiguration3) AddEncryptionFormat(value string) {
	a.EncryptionFormat = append(a.EncryptionFormat, (*EncryptionFormat1Code)(&value))
}
