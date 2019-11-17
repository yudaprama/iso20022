package model

// Choice of format for the pending reason.
type PendingReason26Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason10Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason26Choice) SetCode(value string) {
	p.Code = (*PendingReason10Code)(&value)
}

func (p *PendingReason26Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
