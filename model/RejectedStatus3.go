package model

// Status is rejected.
type RejectedStatus3 struct {

	// Reason for a rejected status in the report.
	Reason []*RejectedStatusReason6 `xml:"Rsn"`

	// Proprietary identification of a reason for a rejected status in the report.
	DataSourceScheme []*GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus3) AddReason() *RejectedStatusReason6 {
	newValue := new(RejectedStatusReason6)
	r.Reason = append(r.Reason, newValue)
	return newValue
}

func (r *RejectedStatus3) AddDataSourceScheme() *GenericIdentification1 {
	newValue := new(GenericIdentification1)
	r.DataSourceScheme = append(r.DataSourceScheme, newValue)
	return newValue
}
