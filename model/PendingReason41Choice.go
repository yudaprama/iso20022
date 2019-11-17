package model

// Choice of format for the pending reason.
type PendingReason41Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason41Choice) SetCode(value string) {
	p.Code = (*PendingReason2Code)(&value)
}

func (p *PendingReason41Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
