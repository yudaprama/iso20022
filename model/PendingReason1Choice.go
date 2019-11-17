package model

// Choice of format for the pending reason.
type PendingReason1Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason1Choice) SetCode(value string) {
	p.Code = (*PendingReason1Code)(&value)
}

func (p *PendingReason1Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
