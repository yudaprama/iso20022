package model

// Reason for the rejected status.
type RejectedStatus8Choice struct {

	// Reason for the rejected status.
	Reason *TransferRejectedStatusReason2Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for a rejected status.
	DataSourceScheme []*GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus8Choice) SetReason(value string) {
	r.Reason = (*TransferRejectedStatusReason2Code)(&value)
}

func (r *RejectedStatus8Choice) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

func (r *RejectedStatus8Choice) AddDataSourceScheme() *GenericIdentification1 {
	newValue := new(GenericIdentification1)
	r.DataSourceScheme = append(r.DataSourceScheme, newValue)
	return newValue
}
