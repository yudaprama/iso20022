package model

// Information about a custody statement of holdings.
type CustodyStatementOfHoldings1 struct {

	// Pagination of the message.
	MessagePagination *Pagination `xml:"MsgPgntn"`

	// General information related to the custody statement of holdings that is being cancelled.
	StatementGeneralDetails *Statement3 `xml:"StmtGnlDtls,omitempty"`

	// Safekeeping or investment account of the statement that is being cancelled.
	AccountDetails *SafekeepingAccount1 `xml:"AcctDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*AggregateBalanceInformation1 `xml:"BalForAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*SubAccountIdentification1 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (c *CustodyStatementOfHoldings1) AddMessagePagination() *Pagination {
	c.MessagePagination = new(Pagination)
	return c.MessagePagination
}

func (c *CustodyStatementOfHoldings1) AddStatementGeneralDetails() *Statement3 {
	c.StatementGeneralDetails = new(Statement3)
	return c.StatementGeneralDetails
}

func (c *CustodyStatementOfHoldings1) AddAccountDetails() *SafekeepingAccount1 {
	c.AccountDetails = new(SafekeepingAccount1)
	return c.AccountDetails
}

func (c *CustodyStatementOfHoldings1) AddBalanceForAccount() *AggregateBalanceInformation1 {
	newValue := new(AggregateBalanceInformation1)
	c.BalanceForAccount = append(c.BalanceForAccount, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings1) AddSubAccountDetails() *SubAccountIdentification1 {
	newValue := new(SubAccountIdentification1)
	c.SubAccountDetails = append(c.SubAccountDetails, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings1) AddTotalValues() *TotalValueInPageAndStatement {
	c.TotalValues = new(TotalValueInPageAndStatement)
	return c.TotalValues
}

func (c *CustodyStatementOfHoldings1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
