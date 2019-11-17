package model

// Unique and unambiguous identifier of a financial institution or a branch of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type BranchAndFinancialInstitutionIdentification3 struct {

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification5Choice `xml:"FinInstnId"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData `xml:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification3) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification5Choice {
	b.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification5Choice)
	return b.FinancialInstitutionIdentification
}

func (b *BranchAndFinancialInstitutionIdentification3) AddBranchIdentification() *BranchData {
	b.BranchIdentification = new(BranchData)
	return b.BranchIdentification
}
