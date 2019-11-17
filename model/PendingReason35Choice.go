package model

// Choice between a standard code and a proprietary code to specify the reason why the instruction/event has a pending status.
type PendingReason35Choice struct {

	// Standard code to specify the reason why the instruction/event has a pending status.
	Code *PendingReason13Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/event has a pending status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason35Choice) SetCode(value string) {
	p.Code = (*PendingReason13Code)(&value)
}

func (p *PendingReason35Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
