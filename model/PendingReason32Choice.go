package model

// Choice between a standard code and a proprietary code to specify the reason why the instruction/event has a pending status.
type PendingReason32Choice struct {

	// Standard code to specify the reason why the instruction/event has a pending status.
	Code *PendingReason4Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/event has a pending status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason32Choice) SetCode(value string) {
	p.Code = (*PendingReason4Code)(&value)
}

func (p *PendingReason32Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
