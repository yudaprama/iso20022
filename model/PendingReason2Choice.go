package model

// Choice of format for the pending reason.
type PendingReason2Choice struct {

	// Specifies the reason why the cancellation request is pending.
	Code *PendingReason6Code `xml:"Cd"`

	// Specifies the reason why the cancellation request is pending.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason2Choice) SetCode(value string) {
	p.Code = (*PendingReason6Code)(&value)
}

func (p *PendingReason2Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
