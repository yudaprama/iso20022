package model

// Choice of format for the pending cancellation reason.
type PendingCancellationReasons2Choice struct {

	// Specifies the reason why the cancellation request is pending.
	Code *PendingReason7Code `xml:"Cd"`

	// Specifies the reason why the cancellation request is pending.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingCancellationReasons2Choice) SetCode(value string) {
	p.Code = (*PendingReason7Code)(&value)
}

func (p *PendingCancellationReasons2Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
