package model

// Identification of the account expressed with an account number and a code.
type AccountIdentificationAndPurpose struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Specifies the purpose of the account.
	Purpose *SecuritiesAccountPurposeType1Code `xml:"Purp"`
}

func (a *AccountIdentificationAndPurpose) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *AccountIdentificationAndPurpose) SetPurpose(value string) {
	a.Purpose = (*SecuritiesAccountPurposeType1Code)(&value)
}
