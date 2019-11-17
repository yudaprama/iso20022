package model

// Information of a single instalment related to an invoice settlement.
type Instalment1 struct {

	// Specifies the progressive number of the single instalment related to an invoice.
	SequenceIdentification *Max70Text `xml:"SeqId"`

	// Due date for the payment of the  invoice instalment.
	PaymentDueDate *ISODate `xml:"PmtDueDt"`

	// Amount of a single instalment related to an invoice.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (i *Instalment1) SetSequenceIdentification(value string) {
	i.SequenceIdentification = (*Max70Text)(&value)
}

func (i *Instalment1) SetPaymentDueDate(value string) {
	i.PaymentDueDate = (*ISODate)(&value)
}

func (i *Instalment1) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}
