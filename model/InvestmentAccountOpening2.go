package model

// Information about the type of opening instruction and identification of the application request.
type InvestmentAccountOpening2 struct {

	// Specifies if the account opening instruction is about a newly created account or a supplementary account.
	OpeningType *AccountOpeningType1Code `xml:"OpngTp"`

	// Unique and unambiguous identifier of the account opening request at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous identification of a transaction, for example, a transfer, as assigned by the investor or account owner.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`
}

func (i *InvestmentAccountOpening2) SetOpeningType(value string) {
	i.OpeningType = (*AccountOpeningType1Code)(&value)
}

func (i *InvestmentAccountOpening2) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening2) AddCounterpartyReference() *AdditionalReference6 {
	i.CounterpartyReference = new(AdditionalReference6)
	return i.CounterpartyReference
}

func (i *InvestmentAccountOpening2) SetExistingAccountIdentification(value string) {
	i.ExistingAccountIdentification = (*Max35Text)(&value)
}
