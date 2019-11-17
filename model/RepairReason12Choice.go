package model

// Choice of format for the repair reason.
type RepairReason12Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason5Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepairReason12Choice) SetCode(value string) {
	r.Code = (*RepairReason5Code)(&value)
}

func (r *RepairReason12Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
