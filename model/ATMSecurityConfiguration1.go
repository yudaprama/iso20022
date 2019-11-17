package model

// Configuration parameters in use by the security device.
type ATMSecurityConfiguration1 struct {

	// Configuration of the cryptographic keys.
	Keys *ATMSecurityConfiguration2 `xml:"Keys,omitempty"`

	// Configuration of the encryption or digital envelope, if the security module is able to perform encryption.
	Encryption *ATMSecurityConfiguration3 `xml:"Ncrptn,omitempty"`

	// MAC (Message Authentication Code) algorithm the security module is able to manage.
	MACAlgorithm []*Algorithm12Code `xml:"MACAlgo,omitempty"`

	// Digest algorithm the security module is able to manage.
	DigestAlgorithm []*Algorithm11Code `xml:"DgstAlgo,omitempty"`

	// Configuration of the digital signatures if the security module is able to perform digital signatures with an asymmetric key.
	DigitalSignature *ATMSecurityConfiguration4 `xml:"DgtlSgntr,omitempty"`

	// Configuration of the PIN online verification.
	PIN *ATMSecurityConfiguration5 `xml:"PIN,omitempty"`

	// Mechanism used to protect the message of the ATM protocol.
	MessageProtection []*MessageProtection1Code `xml:"MsgPrtcn,omitempty"`
}

func (a *ATMSecurityConfiguration1) AddKeys() *ATMSecurityConfiguration2 {
	a.Keys = new(ATMSecurityConfiguration2)
	return a.Keys
}

func (a *ATMSecurityConfiguration1) AddEncryption() *ATMSecurityConfiguration3 {
	a.Encryption = new(ATMSecurityConfiguration3)
	return a.Encryption
}

func (a *ATMSecurityConfiguration1) AddMACAlgorithm(value string) {
	a.MACAlgorithm = append(a.MACAlgorithm, (*Algorithm12Code)(&value))
}

func (a *ATMSecurityConfiguration1) AddDigestAlgorithm(value string) {
	a.DigestAlgorithm = append(a.DigestAlgorithm, (*Algorithm11Code)(&value))
}

func (a *ATMSecurityConfiguration1) AddDigitalSignature() *ATMSecurityConfiguration4 {
	a.DigitalSignature = new(ATMSecurityConfiguration4)
	return a.DigitalSignature
}

func (a *ATMSecurityConfiguration1) AddPIN() *ATMSecurityConfiguration5 {
	a.PIN = new(ATMSecurityConfiguration5)
	return a.PIN
}

func (a *ATMSecurityConfiguration1) AddMessageProtection(value string) {
	a.MessageProtection = append(a.MessageProtection, (*MessageProtection1Code)(&value))
}
