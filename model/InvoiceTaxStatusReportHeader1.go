package model

// Provides the Invoice tax status report header details.
type InvoiceTaxStatusReportHeader1 struct {

	// Party to which the TaxReport is delivered. This message block contains party details for a specific tax authority.
	TaxAuthority *TaxOrganisationIdentification1 `xml:"TaxAuthrty,omitempty"`

	// Identifies the InvoiceTaxReportStatusAdvice message.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Reference to the identification of the InvoiceTaxReport message.
	OriginalMessageIdentification *MessageIdentification1 `xml:"OrgnlMsgId"`

	// Provides the status for the full report.
	ReportStatus *TaxReportingStatus1Code `xml:"RptSts"`

	// Provides the details of the rule which could not be validated.
	ValidationRule []*GenericValidationRuleIdentification1 `xml:"VldtnRule,omitempty"`
}

func (i *InvoiceTaxStatusReportHeader1) AddTaxAuthority() *TaxOrganisationIdentification1 {
	i.TaxAuthority = new(TaxOrganisationIdentification1)
	return i.TaxAuthority
}

func (i *InvoiceTaxStatusReportHeader1) AddMessageIdentification() *MessageIdentification1 {
	i.MessageIdentification = new(MessageIdentification1)
	return i.MessageIdentification
}

func (i *InvoiceTaxStatusReportHeader1) AddOriginalMessageIdentification() *MessageIdentification1 {
	i.OriginalMessageIdentification = new(MessageIdentification1)
	return i.OriginalMessageIdentification
}

func (i *InvoiceTaxStatusReportHeader1) SetReportStatus(value string) {
	i.ReportStatus = (*TaxReportingStatus1Code)(&value)
}

func (i *InvoiceTaxStatusReportHeader1) AddValidationRule() *GenericValidationRuleIdentification1 {
	newValue := new(GenericValidationRuleIdentification1)
	i.ValidationRule = append(i.ValidationRule, newValue)
	return newValue
}
