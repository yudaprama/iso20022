package model

// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
type PaymentInstrument7Choice struct {

	// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
	CreditTransferDetails *CreditTransfer3 `xml:"CdtTrfDtls"`

	// Written order on which instructions are given to an account holder (a financial institution) to pay a stated sum to a named recipient (the payee).
	ChequeDetails *Cheque3 `xml:"ChqDtls"`

	// Part of the investment account to or from which cash entries are made.
	AccountDetails *InvestmentAccount15 `xml:"AcctDtls"`
}

func (p *PaymentInstrument7Choice) AddCreditTransferDetails() *CreditTransfer3 {
	p.CreditTransferDetails = new(CreditTransfer3)
	return p.CreditTransferDetails
}

func (p *PaymentInstrument7Choice) AddChequeDetails() *Cheque3 {
	p.ChequeDetails = new(Cheque3)
	return p.ChequeDetails
}

func (p *PaymentInstrument7Choice) AddAccountDetails() *InvestmentAccount15 {
	p.AccountDetails = new(InvestmentAccount15)
	return p.AccountDetails
}
