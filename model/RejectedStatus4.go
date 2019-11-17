package model

// Status is rejected.
type RejectedStatus4 struct {

	// Reason for a rejected status.
	Reason *RejectedStatusReason4 `xml:"Rsn"`

	// Proprietary identification for a reason of a rejected status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus4) AddReason() *RejectedStatusReason4 {
	r.Reason = new(RejectedStatusReason4)
	return r.Reason
}

func (r *RejectedStatus4) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}
