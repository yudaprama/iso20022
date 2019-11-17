package model

// Choice of format for the pending reason.
type PendingReason15Choice struct {

	// Specifies the reason why the cancellation request is pending.
	Code *PendingReason9Code `xml:"Cd"`

	// Specifies the reason why the cancellation request is pending.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason15Choice) SetCode(value string) {
	p.Code = (*PendingReason9Code)(&value)
}

func (p *PendingReason15Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
