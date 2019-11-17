package model

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification31Choice struct {

	// International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.
	IBAN *IBANIdentifier `xml:"IBAN"`

	// Basic Bank Account Number (BBAN) - identifier used nationally by financial institutions, ie, in individual countries, generally as part of a National Account Numbering Scheme(s), to uniquely identify the account of a customer.
	BBAN *BBANIdentifier `xml:"BBAN"`

	// Universal Payment Identification Code (UPIC) - identifier used by the New York Clearing House to mask confidential data, such as bank accounts and bank routing numbers. UPIC numbers remain with business customers, regardless of banking relationship changes.
	UPIC *UPICIdentifier `xml:"UPIC"`

	// Account number used by financial institutions in individual countries to identify an account of a customer, but not necessarily the bank and branch of the financial institution in which the account is held.
	DomesticAccount *SimpleIdentificationInformation4 `xml:"DmstAcct"`
}

func (a *AccountIdentification31Choice) SetIBAN(value string) {
	a.IBAN = (*IBANIdentifier)(&value)
}

func (a *AccountIdentification31Choice) SetBBAN(value string) {
	a.BBAN = (*BBANIdentifier)(&value)
}

func (a *AccountIdentification31Choice) SetUPIC(value string) {
	a.UPIC = (*UPICIdentifier)(&value)
}

func (a *AccountIdentification31Choice) AddDomesticAccount() *SimpleIdentificationInformation4 {
	a.DomesticAccount = new(SimpleIdentificationInformation4)
	return a.DomesticAccount
}
