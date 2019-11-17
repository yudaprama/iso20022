package model

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification5Choice struct {

	// International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.
	IBAN *IBANIdentifier `xml:"IBAN"`

	// Basic Bank Account Number (BBAN) - identifier used nationally by financial institutions, ie, in individual countries, generally as part of a National Account Numbering Scheme(s), to uniquely identify the account of a customer.
	BBAN *BBANIdentifier `xml:"BBAN"`

	// Account number used by financial institutions in individual countries to identify an account of a customer, but not necessarily the bank and branch of the financial institution in which the account is held.
	DomesticAccount *Max35Text `xml:"DmstAcct"`

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	DepositoryAccount *Max35Text `xml:"DpstryAcct"`
}

func (a *AccountIdentification5Choice) SetIBAN(value string) {
	a.IBAN = (*IBANIdentifier)(&value)
}

func (a *AccountIdentification5Choice) SetBBAN(value string) {
	a.BBAN = (*BBANIdentifier)(&value)
}

func (a *AccountIdentification5Choice) SetDomesticAccount(value string) {
	a.DomesticAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification5Choice) SetDepositoryAccount(value string) {
	a.DepositoryAccount = (*Max35Text)(&value)
}
