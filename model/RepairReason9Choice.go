package model

// Choice of format for the repair reason.
type RepairReason9Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason7Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *RepairReason9Choice) SetCode(value string) {
	r.Code = (*RepairReason7Code)(&value)
}

func (r *RepairReason9Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
