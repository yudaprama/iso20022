package model

// Identifier of an account, as assigned by the account servicer.
type AccountIdentification30Choice struct {

	// PAN of the card identifying the account.
	Card *Min8Max28NumericText `xml:"Card"`

	// Mobile Subscriber Integrated Service Digital Network (i.e. mobile phone number of the SIM card).
	MSISDN *Max16Text `xml:"MSISDN"`

	// E-mail identifying the account.
	EMail *Max35Text `xml:"EMail"`

	// International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.
	IBAN *IBANIdentifier `xml:"IBAN"`

	// Basic Bank Account Number (BBAN) - identifier used nationally by financial institutions, that is, in individual countries, generally as part of a National Account Numbering Scheme(s), to uniquely identify the account of a customer.
	BBAN *BBANIdentifier `xml:"BBAN"`

	// Universal Payment Identification Code (UPIC) - identifier used by the New York Clearing House to mask confidential data, such as bank accounts and bank routing numbers. UPIC numbers remain with business customers, regardless of banking relationship changes.
	UPIC *UPICIdentifier `xml:"UPIC"`

	// Account number used by financial institutions in individual countries to identify an account of a customer, but not necessarily the bank and branch of the financial institution in which the account is held.
	Domestic *Max35Text `xml:"Dmst"`

	// Other identifier.
	Other *Max35Text `xml:"Othr"`
}

func (a *AccountIdentification30Choice) SetCard(value string) {
	a.Card = (*Min8Max28NumericText)(&value)
}

func (a *AccountIdentification30Choice) SetMSISDN(value string) {
	a.MSISDN = (*Max16Text)(&value)
}

func (a *AccountIdentification30Choice) SetEMail(value string) {
	a.EMail = (*Max35Text)(&value)
}

func (a *AccountIdentification30Choice) SetIBAN(value string) {
	a.IBAN = (*IBANIdentifier)(&value)
}

func (a *AccountIdentification30Choice) SetBBAN(value string) {
	a.BBAN = (*BBANIdentifier)(&value)
}

func (a *AccountIdentification30Choice) SetUPIC(value string) {
	a.UPIC = (*UPICIdentifier)(&value)
}

func (a *AccountIdentification30Choice) SetDomestic(value string) {
	a.Domestic = (*Max35Text)(&value)
}

func (a *AccountIdentification30Choice) SetOther(value string) {
	a.Other = (*Max35Text)(&value)
}
