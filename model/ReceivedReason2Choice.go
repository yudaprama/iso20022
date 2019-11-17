package model

// Choice of formats for the reason of a received status.
type ReceivedReason2Choice struct {

	// Reason for the received status expressed as an ISO 20022 code.
	Code *ExternalReceivedReason1Code `xml:"Cd"`

	// Reason for the received status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (r *ReceivedReason2Choice) SetCode(value string) {
	r.Code = (*ExternalReceivedReason1Code)(&value)
}

func (r *ReceivedReason2Choice) AddProprietary() *GenericIdentification36 {
	r.Proprietary = new(GenericIdentification36)
	return r.Proprietary
}
