package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a accepted status.
type AcceptedReason11Choice struct {

	// Standard code to specify additional information about the processed instruction.
	Code *AcknowledgementReason7Code `xml:"Cd"`

	// Proprietary identification of additional information about the processed instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AcceptedReason11Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason7Code)(&value)
}

func (a *AcceptedReason11Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
