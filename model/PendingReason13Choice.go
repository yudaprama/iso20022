package model

// Choice of format for the pending reason.
type PendingReason13Choice struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason13Choice) SetCode(value string) {
	p.Code = (*PendingReason2Code)(&value)
}

func (p *PendingReason13Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
