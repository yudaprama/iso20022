package model

// Specifies the reason for the reversal of the transaction.
type ReversalReason4Choice struct {

	// Reason for the reversal, as published in an external reason code list.
	Code *ExternalReversalReason1Code `xml:"Cd"`

	// Reason for the reversal, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *ReversalReason4Choice) SetCode(value string) {
	r.Code = (*ExternalReversalReason1Code)(&value)
}

func (r *ReversalReason4Choice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
