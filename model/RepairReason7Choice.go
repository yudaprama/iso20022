package model

// Choice of format for the repair reason.
type RepairReason7Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason6Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepairReason7Choice) SetCode(value string) {
	r.Code = (*RepairReason6Code)(&value)
}

func (r *RepairReason7Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
