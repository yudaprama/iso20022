package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a accepted status.
type AcceptedReason1Choice struct {

	// Standard code to specify additional information about the processed instruction.
	Code *AcknowledgementReason4Code `xml:"Cd"`

	// Proprietary identification of additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcceptedReason1Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason4Code)(&value)
}

func (a *AcceptedReason1Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
