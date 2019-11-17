package model

// Provide information about the type of opening instruction and identification of the application request.
type InvestmentAccountOpeningDetails struct {

	// Specifies if the account opening instruction is about a newly created account or a supplementary account.
	OpeningType *AccountOpeningType1Code `xml:"OpngTp"`

	// Unique and unambiguous identifier of the account opening request at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`
}

func (i *InvestmentAccountOpeningDetails) SetOpeningType(value string) {
	i.OpeningType = (*AccountOpeningType1Code)(&value)
}

func (i *InvestmentAccountOpeningDetails) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}
