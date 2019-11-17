package model

// Status is rejected.
type RejectedStatus6 struct {

	// Reason for the rejected status.
	Reason *RejectedStatusReason7Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the rejected status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the rejected status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatus6) SetReason(value string) {
	r.Reason = (*RejectedStatusReason7Code)(&value)
}

func (r *RejectedStatus6) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

func (r *RejectedStatus6) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}

func (r *RejectedStatus6) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}
