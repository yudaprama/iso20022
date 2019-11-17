package model

// Choice of format for the pending reason.
type PendingReason29Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason8Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason29Choice) SetCode(value string) {
	p.Code = (*PendingReason8Code)(&value)
}

func (p *PendingReason29Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
