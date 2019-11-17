package model

// Choice of format for the pending reason.
type PendingReason30Choice struct {

	// Specifies the reason why the cancellation request is pending.
	Code *PendingReason9Code `xml:"Cd"`

	// Specifies the reason why the cancellation request is pending.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason30Choice) SetCode(value string) {
	p.Code = (*PendingReason9Code)(&value)
}

func (p *PendingReason30Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
