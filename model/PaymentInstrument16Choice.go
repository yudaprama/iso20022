package model

// Choice of payment instruments.
type PaymentInstrument16Choice struct {

	// Cash account to credit for the payment of the dividends or of the redeemed investments funds or of interest.
	CashAccountDetails []*CashAccount26 `xml:"CshAcctDtls"`

	// Settlement instructions for a payment by cheque.
	ChequeDetails *Cheque4 `xml:"ChqDtls"`

	// Settlement instructions for a payment by draft.
	BankersDraftDetails *Cheque4 `xml:"BkrsDrftDtls"`
}

func (p *PaymentInstrument16Choice) AddCashAccountDetails() *CashAccount26 {
	newValue := new(CashAccount26)
	p.CashAccountDetails = append(p.CashAccountDetails, newValue)
	return newValue
}

func (p *PaymentInstrument16Choice) AddChequeDetails() *Cheque4 {
	p.ChequeDetails = new(Cheque4)
	return p.ChequeDetails
}

func (p *PaymentInstrument16Choice) AddBankersDraftDetails() *Cheque4 {
	p.BankersDraftDetails = new(Cheque4)
	return p.BankersDraftDetails
}
