package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction23 struct {

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument12Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction23) AddPaymentInstrument() *PaymentInstrument12Choice {
	p.PaymentInstrument = new(PaymentInstrument12Choice)
	return p.PaymentInstrument
}
