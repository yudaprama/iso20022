package model

// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
type PaymentInstrument11Choice struct {

	// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
	CreditTransferDetails *CreditTransfer6 `xml:"CdtTrfDtls"`

	// Written order on which instructions are given to an account holder (a financial institution) to pay a stated sum to a named recipient (the payee).
	ChequeDetails *Cheque3 `xml:"ChqDtls"`

	// Cheque drawn by a bank on itself or its agent. A person who owes money to another buys the draft from a bank for cash and hands it to the creditor.
	BankersDraftDetails *Cheque3 `xml:"BkrsDrftDtls"`

	// Part of the investment account to or from which cash entries are made.
	CashAccountDetails *InvestmentAccount20 `xml:"CshAcctDtls"`
}

func (p *PaymentInstrument11Choice) AddCreditTransferDetails() *CreditTransfer6 {
	p.CreditTransferDetails = new(CreditTransfer6)
	return p.CreditTransferDetails
}

func (p *PaymentInstrument11Choice) AddChequeDetails() *Cheque3 {
	p.ChequeDetails = new(Cheque3)
	return p.ChequeDetails
}

func (p *PaymentInstrument11Choice) AddBankersDraftDetails() *Cheque3 {
	p.BankersDraftDetails = new(Cheque3)
	return p.BankersDraftDetails
}

func (p *PaymentInstrument11Choice) AddCashAccountDetails() *InvestmentAccount20 {
	p.CashAccountDetails = new(InvestmentAccount20)
	return p.CashAccountDetails
}
