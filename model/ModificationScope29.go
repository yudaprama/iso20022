package model

// Scope of the modification to be applied on an identified set of information.
type ModificationScope29 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails *FinancialInstrument51 `xml:"FinInstrmDtls"`
}

func (m *ModificationScope29) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope29) AddFinancialInstrumentDetails() *FinancialInstrument51 {
	m.FinancialInstrumentDetails = new(FinancialInstrument51)
	return m.FinancialInstrumentDetails
}
