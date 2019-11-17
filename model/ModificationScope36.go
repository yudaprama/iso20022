package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope36 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails *FinancialInstrument56 `xml:"FinInstrmDtls"`
}

func (m *ModificationScope36) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope36) AddFinancialInstrumentDetails() *FinancialInstrument56 {
	m.FinancialInstrumentDetails = new(FinancialInstrument56)
	return m.FinancialInstrumentDetails
}
