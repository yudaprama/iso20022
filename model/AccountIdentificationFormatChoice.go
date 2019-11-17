package model

// Choice between formats for account identification.
type AccountIdentificationFormatChoice struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	SimpleIdentification *AccountIdentification1 `xml:"SmplId"`

	// Identification of the account expressed with an account number and a code.
	IdentificationAndPurpose *AccountIdentificationAndPurpose `xml:"IdAndPurp"`

	// Identification of the account expressed with a data source scheme, a code used within the data source scheme and the account identification.
	IdentificationAsDSS *AccountIdentification3 `xml:"IdAsDSS"`
}

func (a *AccountIdentificationFormatChoice) AddSimpleIdentification() *AccountIdentification1 {
	a.SimpleIdentification = new(AccountIdentification1)
	return a.SimpleIdentification
}

func (a *AccountIdentificationFormatChoice) AddIdentificationAndPurpose() *AccountIdentificationAndPurpose {
	a.IdentificationAndPurpose = new(AccountIdentificationAndPurpose)
	return a.IdentificationAndPurpose
}

func (a *AccountIdentificationFormatChoice) AddIdentificationAsDSS() *AccountIdentification3 {
	a.IdentificationAsDSS = new(AccountIdentification3)
	return a.IdentificationAsDSS
}
