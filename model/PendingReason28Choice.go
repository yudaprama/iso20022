package model

// Choice of format for the pending reason.
type PendingReason28Choice struct {

	// Specifies the reason code why the instruction or request is pending.
	Code *PendingReason6Code `xml:"Cd"`

	// Specifies the reason code why the instruction or request is pending, using a proprietary format.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason28Choice) SetCode(value string) {
	p.Code = (*PendingReason6Code)(&value)
}

func (p *PendingReason28Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
