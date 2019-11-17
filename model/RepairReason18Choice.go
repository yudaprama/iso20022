package model

// Choice of format for the repair reason.
type RepairReason18Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason6Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RepairReason18Choice) SetCode(value string) {
	r.Code = (*RepairReason6Code)(&value)
}

func (r *RepairReason18Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
