package model

// Identification of a financial institution.
type FinancialInstitutionIdentification10 struct {

	// Unique identification of the party.
	Party *FinancialInstitutionIdentification8Choice `xml:"Pty"`

	// Legal entity identification as an alternate identification for the party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (f *FinancialInstitutionIdentification10) AddParty() *FinancialInstitutionIdentification8Choice {
	f.Party = new(FinancialInstitutionIdentification8Choice)
	return f.Party
}

func (f *FinancialInstitutionIdentification10) SetLEI(value string) {
	f.LEI = (*LEIIdentifier)(&value)
}
