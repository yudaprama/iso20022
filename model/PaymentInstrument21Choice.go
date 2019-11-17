package model

// Choice between types of payment instrument, that is, cheque, credit transfer or investment account.
type PaymentInstrument21Choice struct {

	// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
	CreditTransferDetails *CreditTransfer8 `xml:"CdtTrfDtls"`

	// Written order on which instructions are given to an account holder (a financial institution) to pay a stated sum to a named recipient (the payee).
	ChequeDetails *Cheque9 `xml:"ChqDtls"`

	// Cheque drawn by a bank on itself or its agent. A person who owes money to another buys the draft from a bank for cash and hands it to the creditor.
	BankersDraftDetails *Cheque9 `xml:"BkrsDrftDtls"`

	// Part of the investment account to or from which cash entries are made.
	CashAccountDetails *InvestmentAccount60 `xml:"CshAcctDtls"`
}

func (p *PaymentInstrument21Choice) AddCreditTransferDetails() *CreditTransfer8 {
	p.CreditTransferDetails = new(CreditTransfer8)
	return p.CreditTransferDetails
}

func (p *PaymentInstrument21Choice) AddChequeDetails() *Cheque9 {
	p.ChequeDetails = new(Cheque9)
	return p.ChequeDetails
}

func (p *PaymentInstrument21Choice) AddBankersDraftDetails() *Cheque9 {
	p.BankersDraftDetails = new(Cheque9)
	return p.BankersDraftDetails
}

func (p *PaymentInstrument21Choice) AddCashAccountDetails() *InvestmentAccount60 {
	p.CashAccountDetails = new(InvestmentAccount60)
	return p.CashAccountDetails
}
