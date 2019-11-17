package model

// Status is rejected.
type RejectedStatus7 struct {

	// Reason for the rejected status.
	Reason *RejectedStatusReason8Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the rejcted status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus7) SetReason(value string) {
	r.Reason = (*RejectedStatusReason8Code)(&value)
}

func (r *RejectedStatus7) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

func (r *RejectedStatus7) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}
