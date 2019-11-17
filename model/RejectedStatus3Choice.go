package model

// Status is rejected.
type RejectedStatus3Choice struct {

	// Reason for a rejected status in the report.
	Reason []*RejectedStatusReason5 `xml:"Rsn"`

	// Proprietary identification of a reason for a rejected status in the report.
	DataSourceScheme []*GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus3Choice) AddReason() *RejectedStatusReason5 {
	newValue := new(RejectedStatusReason5)
	r.Reason = append(r.Reason, newValue)
	return newValue
}

func (r *RejectedStatus3Choice) AddDataSourceScheme() *GenericIdentification1 {
	newValue := new(GenericIdentification1)
	r.DataSourceScheme = append(r.DataSourceScheme, newValue)
	return newValue
}
