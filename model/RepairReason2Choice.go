package model

// Choice of format for the repair reason.
type RepairReason2Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason5Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepairReason2Choice) SetCode(value string) {
	r.Code = (*RepairReason5Code)(&value)
}

func (r *RepairReason2Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
