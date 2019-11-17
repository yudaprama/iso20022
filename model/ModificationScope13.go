package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope13 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Information related to the account to be modified.
	FundDetails *FinancialInstrument29 `xml:"FndDtls"`
}

func (m *ModificationScope13) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope13) AddFundDetails() *FinancialInstrument29 {
	m.FundDetails = new(FinancialInstrument29)
	return m.FundDetails
}
