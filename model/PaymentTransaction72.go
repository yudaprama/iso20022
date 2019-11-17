package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction72 struct {

	// Choice between types of payment instrument, for example, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument21Choice `xml:"PmtInstrm"`
}

func (p *PaymentTransaction72) AddPaymentInstrument() *PaymentInstrument21Choice {
	p.PaymentInstrument = new(PaymentInstrument21Choice)
	return p.PaymentInstrument
}
