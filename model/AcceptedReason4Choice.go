package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a accepted status.
type AcceptedReason4Choice struct {

	// Standard code to specify additional information about the processed instruction.
	Code *AcknowledgementReason8Code `xml:"Cd"`

	// Proprietary identification of additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcceptedReason4Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason8Code)(&value)
}

func (a *AcceptedReason4Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
