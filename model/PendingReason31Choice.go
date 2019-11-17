package model

// Choice of format for the pending reason.
type PendingReason31Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason31Choice) SetCode(value string) {
	p.Code = (*PendingReason1Code)(&value)
}

func (p *PendingReason31Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
