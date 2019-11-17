package model

// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
type PaymentInstrument12Choice struct {

	// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
	PaymentCardDetails *PaymentCard2 `xml:"PmtCardDtls"`

	// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
	CreditTransferDetails *CreditTransfer6 `xml:"CdtTrfDtls"`

	// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
	DirectDebitDetails *DirectDebitMandate4 `xml:"DrctDbtDtls"`

	// Written order on which instructions are given to an account holder (a financial institution) to pay a stated sum to a named recipient (the payee).
	ChequeDetails *Cheque3 `xml:"ChqDtls"`

	// Cheque drawn by a bank on itself or its agent. A person who owes money to another buys the draft from a bank for cash and hands it to the creditor.
	BankersDraftDetails *Cheque3 `xml:"BkrsDrftDtls"`

	// Part of the investment account to or from which cash entries are made.
	CashAccountDetails *InvestmentAccount20 `xml:"CshAcctDtls"`
}

func (p *PaymentInstrument12Choice) AddPaymentCardDetails() *PaymentCard2 {
	p.PaymentCardDetails = new(PaymentCard2)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument12Choice) AddCreditTransferDetails() *CreditTransfer6 {
	p.CreditTransferDetails = new(CreditTransfer6)
	return p.CreditTransferDetails
}

func (p *PaymentInstrument12Choice) AddDirectDebitDetails() *DirectDebitMandate4 {
	p.DirectDebitDetails = new(DirectDebitMandate4)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument12Choice) AddChequeDetails() *Cheque3 {
	p.ChequeDetails = new(Cheque3)
	return p.ChequeDetails
}

func (p *PaymentInstrument12Choice) AddBankersDraftDetails() *Cheque3 {
	p.BankersDraftDetails = new(Cheque3)
	return p.BankersDraftDetails
}

func (p *PaymentInstrument12Choice) AddCashAccountDetails() *InvestmentAccount20 {
	p.CashAccountDetails = new(InvestmentAccount20)
	return p.CashAccountDetails
}
