package model

// Information about a statement of investment fund transactions.
type StatementOfInvestmentFundTransactions3 struct {

	// General information related to the investment fund statement of transactions that is being cancelled.
	StatementGeneralDetails *Statement8 `xml:"StmtGnlDtls,omitempty"`

	// Information related to an investment account of the statement that is being cancelled.
	InvestmentAccountDetails *InvestmentAccount43 `xml:"InvstmtAcctDtls,omitempty"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnAccount []*InvestmentFundTransactionsByFund3 `xml:"TxOnAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*SubAccountIdentification36 `xml:"SubAcctDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *StatementOfInvestmentFundTransactions3) AddStatementGeneralDetails() *Statement8 {
	s.StatementGeneralDetails = new(Statement8)
	return s.StatementGeneralDetails
}

func (s *StatementOfInvestmentFundTransactions3) AddInvestmentAccountDetails() *InvestmentAccount43 {
	s.InvestmentAccountDetails = new(InvestmentAccount43)
	return s.InvestmentAccountDetails
}

func (s *StatementOfInvestmentFundTransactions3) AddTransactionOnAccount() *InvestmentFundTransactionsByFund3 {
	newValue := new(InvestmentFundTransactionsByFund3)
	s.TransactionOnAccount = append(s.TransactionOnAccount, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactions3) AddSubAccountDetails() *SubAccountIdentification36 {
	newValue := new(SubAccountIdentification36)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *StatementOfInvestmentFundTransactions3) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
