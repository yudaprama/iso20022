package model

// Choice of format for the pending reason.
type PendingReason36Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason10Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason36Choice) SetCode(value string) {
	p.Code = (*PendingReason10Code)(&value)
}

func (p *PendingReason36Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
