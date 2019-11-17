package model

// Choice of format for the pending reason.
type PendingReason42Choice struct {

	// Specifies the reason why the cancellation request is pending.
	Code *PendingReason9Code `xml:"Cd"`

	// Specifies the reason why the cancellation request is pending.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingReason42Choice) SetCode(value string) {
	p.Code = (*PendingReason9Code)(&value)
}

func (p *PendingReason42Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
