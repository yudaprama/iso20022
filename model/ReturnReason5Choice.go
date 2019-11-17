package model

// Specifies the reason for the return of the transaction.
type ReturnReason5Choice struct {

	// Reason for the return, as published in an external reason code list.
	Code *ExternalReturnReason1Code `xml:"Cd"`

	// Reason for the return, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *ReturnReason5Choice) SetCode(value string) {
	r.Code = (*ExternalReturnReason1Code)(&value)
}

func (r *ReturnReason5Choice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
