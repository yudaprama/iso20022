package model

// Choice of format for the pending cancellation reason.
type PendingCancellationReasons5Choice struct {

	// Specifies the reason why the cancellation request is pending.
	Code *PendingReason7Code `xml:"Cd"`

	// Specifies the reason why the cancellation request is pending.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingCancellationReasons5Choice) SetCode(value string) {
	p.Code = (*PendingReason7Code)(&value)
}

func (p *PendingCancellationReasons5Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
