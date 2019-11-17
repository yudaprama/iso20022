package model

// Choice of format for the pending reason.
type PendingReason3Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason3Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason3Choice) SetCode(value string) {
	p.Code = (*PendingReason3Code)(&value)
}

func (p *PendingReason3Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
