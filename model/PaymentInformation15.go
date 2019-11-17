package model

// Payment terms and conditions related to a single invoice to be financed.
type PaymentInformation15 struct {

	// Payment method that will be used for invoice payment to transfer the funds to the creditor.
	PaymentMethod *PaymentMethod4Code `xml:"PmtMtd"`

	// Unambiguous identification of the account used for payment settlement.
	PaymentAccount *CashAccount7 `xml:"PmtAcct,omitempty"`
}

func (p *PaymentInformation15) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod4Code)(&value)
}

func (p *PaymentInformation15) AddPaymentAccount() *CashAccount7 {
	p.PaymentAccount = new(CashAccount7)
	return p.PaymentAccount
}
