package model

// Choice of format for the repair reason.
type RepairReason10Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason4Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepairReason10Choice) SetCode(value string) {
	r.Code = (*RepairReason4Code)(&value)
}

func (r *RepairReason10Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
