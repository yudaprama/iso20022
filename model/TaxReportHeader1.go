package model

// Defines message level identification, number of individual tax reports and tax authority.
type TaxReportHeader1 struct {

	// Unique message identification.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Number of TaxReports in this message. Seller can send all TaxReports in the same file.
	NumberOfTaxReports *Number `xml:"NbOfTaxRpts,omitempty"`

	// Party to which the TaxReport is delivered. This message block contains party details for a specific tax authority.
	TaxAuthority []*TaxOrganisationIdentification1 `xml:"TaxAuthrty,omitempty"`
}

func (t *TaxReportHeader1) AddMessageIdentification() *MessageIdentification1 {
	t.MessageIdentification = new(MessageIdentification1)
	return t.MessageIdentification
}

func (t *TaxReportHeader1) SetNumberOfTaxReports(value string) {
	t.NumberOfTaxReports = (*Number)(&value)
}

func (t *TaxReportHeader1) AddTaxAuthority() *TaxOrganisationIdentification1 {
	newValue := new(TaxOrganisationIdentification1)
	t.TaxAuthority = append(t.TaxAuthority, newValue)
	return newValue
}
