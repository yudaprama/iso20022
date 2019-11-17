package model

// Choice of format for the pending reason.
type PendingReason16Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason10Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason16Choice) SetCode(value string) {
	p.Code = (*PendingReason10Code)(&value)
}

func (p *PendingReason16Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
