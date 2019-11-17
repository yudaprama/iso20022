package model

// Reason for which a cancellation is rejected.
type RejectedCancellationStatusReason1Choice struct {

	// Reason for the rejected status.
	Reason *RejectedCancellationStatusReason1Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the rejected status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedCancellationStatusReason1Choice) SetReason(value string) {
	r.Reason = (*RejectedCancellationStatusReason1Code)(&value)
}

func (r *RejectedCancellationStatusReason1Choice) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

func (r *RejectedCancellationStatusReason1Choice) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}
