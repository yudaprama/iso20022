package model

// Unique identifier of an account, as assigned by the account servicer.
type CashAccountIdentification5Choice struct {

	// International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.
	IBAN *IBAN2007Identifier `xml:"IBAN"`

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	Proprietary *Max34Text `xml:"Prtry"`
}

func (c *CashAccountIdentification5Choice) SetIBAN(value string) {
	c.IBAN = (*IBAN2007Identifier)(&value)
}

func (c *CashAccountIdentification5Choice) SetProprietary(value string) {
	c.Proprietary = (*Max34Text)(&value)
}
