package model

// Choice of format for the repair reason.
type RepairReason13Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RepairReason5Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RepairReason13Choice) SetCode(value string) {
	r.Code = (*RepairReason5Code)(&value)
}

func (r *RepairReason13Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
