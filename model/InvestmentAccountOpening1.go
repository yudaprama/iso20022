package model

// Provides information about the type of opening instruction and identification of the application request.
type InvestmentAccountOpening1 struct {

	// Specifies if the account opening instruction is about a newly created account or a supplementary account.
	OpeningType *AccountOpeningType1Code `xml:"OpngTp"`

	// Unique and unambiguous identifier of the account opening request at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer as allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (i *InvestmentAccountOpening1) SetOpeningType(value string) {
	i.OpeningType = (*AccountOpeningType1Code)(&value)
}

func (i *InvestmentAccountOpening1) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening1) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening1) AddCounterpartyReference() *AdditionalReference2 {
	i.CounterpartyReference = new(AdditionalReference2)
	return i.CounterpartyReference
}
