package model

// Choice between types of payment instrument, that is, cheque, credit transfer, direct debit, investment account or payment card.
type PaymentInstrument20Choice struct {

	// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
	PaymentCardDetails *PaymentCard25 `xml:"PmtCardDtls"`

	// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
	CreditTransferDetails *CreditTransfer8 `xml:"CdtTrfDtls"`

	// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
	DirectDebitDetails *DirectDebitMandate6 `xml:"DrctDbtDtls"`

	// Written order on which instructions are given to an account holder (a financial institution) to pay a stated sum to a named recipient (the payee).
	ChequeDetails *Cheque9 `xml:"ChqDtls"`

	// Cheque drawn by a bank on itself or its agent. A person who owes money to another buys the draft from a bank for cash and hands it to the creditor.
	BankersDraftDetails *Cheque9 `xml:"BkrsDrftDtls"`

	// Part of the investment account to or from which cash entries are made.
	CashAccountDetails *InvestmentAccount60 `xml:"CshAcctDtls"`
}

func (p *PaymentInstrument20Choice) AddPaymentCardDetails() *PaymentCard25 {
	p.PaymentCardDetails = new(PaymentCard25)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument20Choice) AddCreditTransferDetails() *CreditTransfer8 {
	p.CreditTransferDetails = new(CreditTransfer8)
	return p.CreditTransferDetails
}

func (p *PaymentInstrument20Choice) AddDirectDebitDetails() *DirectDebitMandate6 {
	p.DirectDebitDetails = new(DirectDebitMandate6)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument20Choice) AddChequeDetails() *Cheque9 {
	p.ChequeDetails = new(Cheque9)
	return p.ChequeDetails
}

func (p *PaymentInstrument20Choice) AddBankersDraftDetails() *Cheque9 {
	p.BankersDraftDetails = new(Cheque9)
	return p.BankersDraftDetails
}

func (p *PaymentInstrument20Choice) AddCashAccountDetails() *InvestmentAccount60 {
	p.CashAccountDetails = new(InvestmentAccount60)
	return p.CashAccountDetails
}
