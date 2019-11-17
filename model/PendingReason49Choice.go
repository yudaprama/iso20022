package model

// Choice between a standard code and a proprietary code to specify the reason why the instruction/event has a pending status.
type PendingReason49Choice struct {

	// Standard code to specify the reason why the instruction/event has a pending status.
	Code *PendingReason14Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/event has a pending status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason49Choice) SetCode(value string) {
	p.Code = (*PendingReason14Code)(&value)
}

func (p *PendingReason49Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
