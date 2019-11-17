package model

// Provides the details of each individual invoice tax report transaction.
type InvoiceTaxReportTransactionStatus1 struct {

	// Report identification, for example invoice number or report number from point of sales system.
	TaxReportIdentification *Max35Text `xml:"TaxRptId"`

	// Defines status of the reported transaction.
	Status *TaxReportingStatus2Code `xml:"Sts"`

	// Provides the details of the rule which could not be validated.
	ValidationRule []*GenericValidationRuleIdentification1 `xml:"VldtnRule,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InvoiceTaxReportTransactionStatus1) SetTaxReportIdentification(value string) {
	i.TaxReportIdentification = (*Max35Text)(&value)
}

func (i *InvoiceTaxReportTransactionStatus1) SetStatus(value string) {
	i.Status = (*TaxReportingStatus2Code)(&value)
}

func (i *InvoiceTaxReportTransactionStatus1) AddValidationRule() *GenericValidationRuleIdentification1 {
	newValue := new(GenericValidationRuleIdentification1)
	i.ValidationRule = append(i.ValidationRule, newValue)
	return newValue
}

func (i *InvoiceTaxReportTransactionStatus1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
