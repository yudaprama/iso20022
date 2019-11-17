package model

// Reason for the return of the transaction.
type ReturnReason1Choice struct {

	// Reason for the return in a coded form.
	Code *TransactionRejectReason2Code `xml:"Cd"`

	// Reason for the return not catered for by the available codes.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *ReturnReason1Choice) SetCode(value string) {
	r.Code = (*TransactionRejectReason2Code)(&value)
}

func (r *ReturnReason1Choice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
