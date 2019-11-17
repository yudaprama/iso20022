package model

// Choice between a standard code and a proprietary code to specify the reason why the instruction/event has a pending status.
type PendingReason22Choice struct {

	// Standard code to specify the reason why the instruction/event has a pending status.
	Code *PendingReason11Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/event has a pending status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason22Choice) SetCode(value string) {
	p.Code = (*PendingReason11Code)(&value)
}

func (p *PendingReason22Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
