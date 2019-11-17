package model

// Choice of format for the pending reason.
type PendingReason47Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason8Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason47Choice) SetCode(value string) {
	p.Code = (*PendingReason8Code)(&value)
}

func (p *PendingReason47Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
