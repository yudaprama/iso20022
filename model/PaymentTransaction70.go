package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction70 struct {

	// Choice between types of payment instrument, for example, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument20Choice `xml:"PmtInstrm"`
}

func (p *PaymentTransaction70) AddPaymentInstrument() *PaymentInstrument20Choice {
	p.PaymentInstrument = new(PaymentInstrument20Choice)
	return p.PaymentInstrument
}
