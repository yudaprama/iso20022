package model

// Organisation established primarily to provide financial services.
type BranchAndFinancialInstitutionIdentification struct {

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification1 `xml:"FinInstnId"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData `xml:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification1 {
	b.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification1)
	return b.FinancialInstitutionIdentification
}

func (b *BranchAndFinancialInstitutionIdentification) AddBranchIdentification() *BranchData {
	b.BranchIdentification = new(BranchData)
	return b.BranchIdentification
}
