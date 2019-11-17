package model

// Choice of format for the pending cancellation reason.
type PendingCancellationReasons4Choice struct {

	// Specifies the reason why the cancellation request is pending.
	Code *PendingReason7Code `xml:"Cd"`

	// Specifies the reason why the cancellation request is pending.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingCancellationReasons4Choice) SetCode(value string) {
	p.Code = (*PendingReason7Code)(&value)
}

func (p *PendingCancellationReasons4Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
