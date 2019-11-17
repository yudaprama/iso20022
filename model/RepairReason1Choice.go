package model

// Choice of format for the repair reason.
type RepairReason1Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason4Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepairReason1Choice) SetCode(value string) {
	r.Code = (*RepairReason4Code)(&value)
}

func (r *RepairReason1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
