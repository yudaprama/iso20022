package model

// Reason for the reversal of the transaction.
type ReversalReason1Choice struct {

	// Reason for the reversal in a coded form.
	Code *TransactionReversalReason1Code `xml:"Cd"`

	// Reason for the reversal not catered for by the available codes.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *ReversalReason1Choice) SetCode(value string) {
	r.Code = (*TransactionReversalReason1Code)(&value)
}

func (r *ReversalReason1Choice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
