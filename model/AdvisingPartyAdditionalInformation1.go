package model

// Additional information related to the advising party.
type AdvisingPartyAdditionalInformation1 struct {

	// Unique and unambiguous identifier assigned as a reference.
	ReferenceNumber *Max35Text `xml:"RefNb,omitempty"`

	// Additional information specific to the bank-to-beneficiary communication.
	BankToBeneficiaryInformation []*Max2000Text `xml:"BkToBnfcryInf,omitempty"`
}

func (a *AdvisingPartyAdditionalInformation1) SetReferenceNumber(value string) {
	a.ReferenceNumber = (*Max35Text)(&value)
}

func (a *AdvisingPartyAdditionalInformation1) AddBankToBeneficiaryInformation(value string) {
	a.BankToBeneficiaryInformation = append(a.BankToBeneficiaryInformation, (*Max2000Text)(&value))
}
