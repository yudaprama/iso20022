package model

// Cash settlement parties and accounts.
type CashSettlement1 struct {

	// Account to credit or debit. When this is an account to debit, this is for the payment of a subscription to an investment fund, a savings plan payment, the purchase of securities or the payment of charges. When this is an account to credit,  this is for the payment of an amount as a result of a redemption of investment fund units, the sale of securities, interest and dividend payments. A single account may be specified for all cash movements on the account or cash accounts may be specified for specific types of transactions on the account.
	CashAccountDetails []*CashAccount33 `xml:"CshAcctDtls,omitempty"`

	// Method of payment other than a cash account.
	OtherCashSettlementDetails []*PaymentInstrument13 `xml:"OthrCshSttlmDtls,omitempty"`
}

func (c *CashSettlement1) AddCashAccountDetails() *CashAccount33 {
	newValue := new(CashAccount33)
	c.CashAccountDetails = append(c.CashAccountDetails, newValue)
	return newValue
}

func (c *CashSettlement1) AddOtherCashSettlementDetails() *PaymentInstrument13 {
	newValue := new(PaymentInstrument13)
	c.OtherCashSettlementDetails = append(c.OtherCashSettlementDetails, newValue)
	return newValue
}
