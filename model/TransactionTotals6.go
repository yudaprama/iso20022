package model

// Additional count which may be utilised for reconciliation.
type TransactionTotals6 struct {

	// Sum number of all authorisation transactions.
	Authorisation *Number `xml:"Authstn,omitempty"`

	// Sum number of all reversed authorisation transactions.
	AuthorisationReversal *Number `xml:"AuthstnRvsl,omitempty"`

	// Sum number of all inquiry transactions.
	Inquiry *Number `xml:"Nqry,omitempty"`

	// Sum number of all reversed inquiry transactions.
	InquiryReversal *Number `xml:"NqryRvsl,omitempty"`

	// Sum number of all financial presentment payment transactions processed.
	Payments *Number `xml:"Pmts,omitempty"`

	// Sum number of all financial presentment payment transactions which have been reversed.
	PaymentReversal *Number `xml:"PmtRvsl,omitempty"`

	// Sum number of all financial presentment transactions processed.
	Transfer *Number `xml:"Trf,omitempty"`

	// Sum number of all reversal transactions processed.
	TransferReversal *Number `xml:"TrfRvsl,omitempty"`

	// Sum number of all fee collection transactions.
	FeeCollection *Number `xml:"FeeColltn,omitempty"`
}

func (t *TransactionTotals6) SetAuthorisation(value string) {
	t.Authorisation = (*Number)(&value)
}

func (t *TransactionTotals6) SetAuthorisationReversal(value string) {
	t.AuthorisationReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetInquiry(value string) {
	t.Inquiry = (*Number)(&value)
}

func (t *TransactionTotals6) SetInquiryReversal(value string) {
	t.InquiryReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetPayments(value string) {
	t.Payments = (*Number)(&value)
}

func (t *TransactionTotals6) SetPaymentReversal(value string) {
	t.PaymentReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetTransfer(value string) {
	t.Transfer = (*Number)(&value)
}

func (t *TransactionTotals6) SetTransferReversal(value string) {
	t.TransferReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetFeeCollection(value string) {
	t.FeeCollection = (*Number)(&value)
}
