package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction22 struct {

	// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument11Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction22) AddPaymentInstrument() *PaymentInstrument11Choice {
	p.PaymentInstrument = new(PaymentInstrument11Choice)
	return p.PaymentInstrument
}
