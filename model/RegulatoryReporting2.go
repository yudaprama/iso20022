package model

// Information needed due to regulatory and/or statutory requirements.
type RegulatoryReporting2 struct {

	// Identifies whether the regulatory reporting information applies to the debit side, to the credit side or to both debit and credit sides of the transaction.
	DebitCreditReportingIndicator *RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`

	// Entity requiring the regulatory reporting information.
	Authority *RegulatoryAuthority `xml:"Authrty,omitempty"`

	// Details related to the regulatory reporting information.
	RegulatoryDetails *StructuredRegulatoryReporting2 `xml:"RgltryDtls,omitempty"`
}

func (r *RegulatoryReporting2) SetDebitCreditReportingIndicator(value string) {
	r.DebitCreditReportingIndicator = (*RegulatoryReportingType1Code)(&value)
}

func (r *RegulatoryReporting2) AddAuthority() *RegulatoryAuthority {
	r.Authority = new(RegulatoryAuthority)
	return r.Authority
}

func (r *RegulatoryReporting2) AddRegulatoryDetails() *StructuredRegulatoryReporting2 {
	r.RegulatoryDetails = new(StructuredRegulatoryReporting2)
	return r.RegulatoryDetails
}
