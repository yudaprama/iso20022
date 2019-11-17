package model

// Choice of format for the pending reason.
type PendingReason37Choice struct {

	// Specifies the reason code why the instruction or request is pending.
	Code *PendingReason6Code `xml:"Cd"`

	// Specifies the reason code why the instruction or request is pending, using a proprietary format.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason37Choice) SetCode(value string) {
	p.Code = (*PendingReason6Code)(&value)
}

func (p *PendingReason37Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
