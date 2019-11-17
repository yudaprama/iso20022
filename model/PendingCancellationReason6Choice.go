package model

// Choice between a standard code or a proprietary code to specify the reason why a cancellation request sent for the related instruction is pending.
type PendingCancellationReason6Choice struct {

	// Standard code to specify the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingCancellationReason5Code `xml:"Cd"`

	// Proprietary identification of the reason why a cancellation request sent for the related instruction is pending.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingCancellationReason6Choice) SetCode(value string) {
	p.Code = (*PendingCancellationReason5Code)(&value)
}

func (p *PendingCancellationReason6Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
