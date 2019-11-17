package model

// Set of elements used to uniquely and unambiguously identify a financial institution or a branch of a financial institution.
type BranchAndFinancialInstitutionIdentification4 struct {

	// Unique and unambiguous identification of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification7 `xml:"FinInstnId"`

	// Identifies a specific branch of a financial institution.
	//
	// Usage: This component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData2 `xml:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification4) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification7 {
	b.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification7)
	return b.FinancialInstitutionIdentification
}

func (b *BranchAndFinancialInstitutionIdentification4) AddBranchIdentification() *BranchData2 {
	b.BranchIdentification = new(BranchData2)
	return b.BranchIdentification
}
