package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a accepted status.
type AcceptedReason3Choice struct {

	// Standard code to specify additional information about the processed instruction.
	Code *AcknowledgementReason7Code `xml:"Cd"`

	// Proprietary identification of additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcceptedReason3Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason7Code)(&value)
}

func (a *AcceptedReason3Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
