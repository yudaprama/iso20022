package model

// Choice of formats for the reason of an accepted status.
type AcceptedReason7Choice struct {

	// Reason for the accepted status expressed as an ISO 20022 code.
	Code *ExternalAcceptedReason1Code `xml:"Cd"`

	// Reason for the accepted status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (a *AcceptedReason7Choice) SetCode(value string) {
	a.Code = (*ExternalAcceptedReason1Code)(&value)
}

func (a *AcceptedReason7Choice) AddProprietary() *GenericIdentification36 {
	a.Proprietary = new(GenericIdentification36)
	return a.Proprietary
}
