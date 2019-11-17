package model

// Choice between a standard code or a proprietary code to specify the reason why a cancellation request sent for the related instruction is pending.
type PendingCancellationReason1Choice struct {

	// Standard code to specify the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingCancellationReason4Code `xml:"Cd"`

	// Proprietary identification of the reason why a cancellation request sent for the related instruction is pending.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingCancellationReason1Choice) SetCode(value string) {
	p.Code = (*PendingCancellationReason4Code)(&value)
}

func (p *PendingCancellationReason1Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
