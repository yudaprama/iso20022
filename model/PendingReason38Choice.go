package model

// Choice of format for the pending reason.
type PendingReason38Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason38Choice) SetCode(value string) {
	p.Code = (*PendingReason1Code)(&value)
}

func (p *PendingReason38Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
