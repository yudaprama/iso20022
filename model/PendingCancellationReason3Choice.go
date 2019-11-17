package model

// Choice between a standard code or a proprietary code to specify the reason why a cancellation request sent for the related instruction is pending.
type PendingCancellationReason3Choice struct {

	// Standard code to specify the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingCancellationReason4Code `xml:"Cd"`

	// Proprietary identification of the reason why a cancellation request sent for the related instruction is pending.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingCancellationReason3Choice) SetCode(value string) {
	p.Code = (*PendingCancellationReason4Code)(&value)
}

func (p *PendingCancellationReason3Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
