package model

// Choice of format for the repair reason.
type RepairReason14Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason4Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RepairReason14Choice) SetCode(value string) {
	r.Code = (*RepairReason4Code)(&value)
}

func (r *RepairReason14Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
