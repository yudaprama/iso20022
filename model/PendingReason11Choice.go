package model

// Choice of format for the pending reason.
type PendingReason11Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason8Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason11Choice) SetCode(value string) {
	p.Code = (*PendingReason8Code)(&value)
}

func (p *PendingReason11Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
