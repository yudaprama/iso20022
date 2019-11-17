package model

// Cryptographic Key component.
type CryptographicKey7 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id,omitempty"`

	// Identification of the security domain.
	SecurityDomainIdentification *Max35Text `xml:"SctyDomnId,omitempty"`

	// Additional identification of the key, for instance to derive the key.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn"`

	// Sequence counter of the cryptographic key.
	SequenceCounter *Number `xml:"SeqCntr,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Value for checking a cryptographic key.
	KeyCheckValue *Max35Binary `xml:"KeyChckVal,omitempty"`

	// Current status of the key.
	CurrentStatus *ATMStatus3Code `xml:"CurSts"`

	// Reason for which the key has been stopped.
	FailureReason *FailureReason6Code `xml:"FailrRsn,omitempty"`
}

func (c *CryptographicKey7) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *CryptographicKey7) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey7) SetSecurityDomainIdentification(value string) {
	c.SecurityDomainIdentification = (*Max35Text)(&value)
}

func (c *CryptographicKey7) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey7) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey7) SetSequenceCounter(value string) {
	c.SequenceCounter = (*Number)(&value)
}

func (c *CryptographicKey7) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey7) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey7) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey7) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey7) SetKeyCheckValue(value string) {
	c.KeyCheckValue = (*Max35Binary)(&value)
}

func (c *CryptographicKey7) SetCurrentStatus(value string) {
	c.CurrentStatus = (*ATMStatus3Code)(&value)
}

func (c *CryptographicKey7) SetFailureReason(value string) {
	c.FailureReason = (*FailureReason6Code)(&value)
}
