package model

// Choice of format for the repair reason.
type RepairReason11Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason6Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepairReason11Choice) SetCode(value string) {
	r.Code = (*RepairReason6Code)(&value)
}

func (r *RepairReason11Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
