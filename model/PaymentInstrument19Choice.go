package model

// Choice of payment instruments.
type PaymentInstrument19Choice struct {

	// Settlement instructions for a payment by cheque.
	ChequeDetails *Cheque4 `xml:"ChqDtls"`

	// Settlement instructions for a payment by draft.
	BankersDraftDetails *Cheque4 `xml:"BkrsDrftDtls"`
}

func (p *PaymentInstrument19Choice) AddChequeDetails() *Cheque4 {
	p.ChequeDetails = new(Cheque4)
	return p.ChequeDetails
}

func (p *PaymentInstrument19Choice) AddBankersDraftDetails() *Cheque4 {
	p.BankersDraftDetails = new(Cheque4)
	return p.BankersDraftDetails
}
