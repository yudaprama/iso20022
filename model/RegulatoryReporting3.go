package model

// Information needed due to regulatory and/or statutory requirements.
type RegulatoryReporting3 struct {

	// Identifies whether the regulatory reporting information applies to the debit side, to the credit side or to both debit and credit sides of the transaction.
	DebitCreditReportingIndicator *RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`

	// Entity requiring the regulatory reporting information.
	Authority *RegulatoryAuthority2 `xml:"Authrty,omitempty"`

	// Set of elements used to provide details on the regulatory reporting information.
	Details []*StructuredRegulatoryReporting3 `xml:"Dtls,omitempty"`
}

func (r *RegulatoryReporting3) SetDebitCreditReportingIndicator(value string) {
	r.DebitCreditReportingIndicator = (*RegulatoryReportingType1Code)(&value)
}

func (r *RegulatoryReporting3) AddAuthority() *RegulatoryAuthority2 {
	r.Authority = new(RegulatoryAuthority2)
	return r.Authority
}

func (r *RegulatoryReporting3) AddDetails() *StructuredRegulatoryReporting3 {
	newValue := new(StructuredRegulatoryReporting3)
	r.Details = append(r.Details, newValue)
	return newValue
}
