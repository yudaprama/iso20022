package model

// Information about a custody statement of holdings.
type CustodyStatementOfHoldings2 struct {

	// General information related to the custody statement of holdings that is being cancelled.
	StatementGeneralDetails *Statement7 `xml:"StmtGnlDtls,omitempty"`

	// Safekeeping or investment account of the statement that is being cancelled.
	AccountDetails *SafekeepingAccount2 `xml:"AcctDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*AggregateBalanceInformation4 `xml:"BalForAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*SubAccountIdentification5 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (c *CustodyStatementOfHoldings2) AddStatementGeneralDetails() *Statement7 {
	c.StatementGeneralDetails = new(Statement7)
	return c.StatementGeneralDetails
}

func (c *CustodyStatementOfHoldings2) AddAccountDetails() *SafekeepingAccount2 {
	c.AccountDetails = new(SafekeepingAccount2)
	return c.AccountDetails
}

func (c *CustodyStatementOfHoldings2) AddBalanceForAccount() *AggregateBalanceInformation4 {
	newValue := new(AggregateBalanceInformation4)
	c.BalanceForAccount = append(c.BalanceForAccount, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings2) AddSubAccountDetails() *SubAccountIdentification5 {
	newValue := new(SubAccountIdentification5)
	c.SubAccountDetails = append(c.SubAccountDetails, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings2) AddTotalValues() *TotalValueInPageAndStatement {
	c.TotalValues = new(TotalValueInPageAndStatement)
	return c.TotalValues
}

func (c *CustodyStatementOfHoldings2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
