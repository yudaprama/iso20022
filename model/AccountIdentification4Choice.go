package model

// Specifies the unique identification of an account as assigned by the account servicer.
type AccountIdentification4Choice struct {

	// International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.
	IBAN *IBAN2007Identifier `xml:"IBAN"`

	// Unique identification of an account, as assigned by the account servicer, using an identification scheme.
	Other *GenericAccountIdentification1 `xml:"Othr"`
}

func (a *AccountIdentification4Choice) SetIBAN(value string) {
	a.IBAN = (*IBAN2007Identifier)(&value)
}

func (a *AccountIdentification4Choice) AddOther() *GenericAccountIdentification1 {
	a.Other = new(GenericAccountIdentification1)
	return a.Other
}
