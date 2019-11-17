package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope12 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Information related to the account to be modified.
	FundDetails *FinancialInstrument10 `xml:"FndDtls"`
}

func (m *ModificationScope12) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope12) AddFundDetails() *FinancialInstrument10 {
	m.FundDetails = new(FinancialInstrument10)
	return m.FundDetails
}
