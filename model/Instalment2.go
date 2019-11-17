package model

// Specifies a single instalment related to an invoice settlement and optional reconciliation information.
// Reconciliation information is used to indicate the amount to be allocated to a particular instalment of a financial document.
type Instalment2 struct {

	// Specifies the progressive number of the single instalment related to an invoice.
	SequenceIdentification *Max70Text `xml:"SeqId"`

	// Due date for the payment of the financing item instalment.
	PaymentDueDate *ISODate `xml:"PmtDueDt"`

	// Amount of a single instalment related to an invoice.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Desired payment instrument to be used for the instalment.
	PaymentInstrument *PaymentMeans1 `xml:"PmtInstrm,omitempty"`
}

func (i *Instalment2) SetSequenceIdentification(value string) {
	i.SequenceIdentification = (*Max70Text)(&value)
}

func (i *Instalment2) SetPaymentDueDate(value string) {
	i.PaymentDueDate = (*ISODate)(&value)
}

func (i *Instalment2) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *Instalment2) AddPaymentInstrument() *PaymentMeans1 {
	i.PaymentInstrument = new(PaymentMeans1)
	return i.PaymentInstrument
}
