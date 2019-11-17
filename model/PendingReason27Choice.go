package model

// Choice of format for the pending reason.
type PendingReason27Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason27Choice) SetCode(value string) {
	p.Code = (*PendingReason2Code)(&value)
}

func (p *PendingReason27Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
