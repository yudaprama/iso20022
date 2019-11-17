package model

// Balance details of a registered contract.
type ContractBalance1 struct {

	// Specifies the type of the contract balance.
	Type *ContractBalanceType1Choice `xml:"Tp"`

	// Currency and amount of money of the contract balance.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the balance is a credit or a debit balance. A zero balance is considered to be a credit balance
	CreditDebitIndicator *CreditDebit3Code `xml:"CdtDbtInd"`
}

func (c *ContractBalance1) AddType() *ContractBalanceType1Choice {
	c.Type = new(ContractBalanceType1Choice)
	return c.Type
}

func (c *ContractBalance1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *ContractBalance1) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebit3Code)(&value)
}
