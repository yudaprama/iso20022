package model

// Reason for which a transaction report is rejected.
type RejectedStatusReason9Choice struct {

	// Reason for the rejected status.
	Reason *RejectedStatusReason9Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the rejected status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatusReason9Choice) SetReason(value string) {
	r.Reason = (*RejectedStatusReason9Code)(&value)
}

func (r *RejectedStatusReason9Choice) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

func (r *RejectedStatusReason9Choice) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}
