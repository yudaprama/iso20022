package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccountIdentification4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnSubAccount []*InvestmentFundTransactionsByFund1 `xml:"TxOnSubAcct,omitempty"`
}

func (s *SubAccountIdentification4) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification4) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification4) AddTransactionOnSubAccount() *InvestmentFundTransactionsByFund1 {
	newValue := new(InvestmentFundTransactionsByFund1)
	s.TransactionOnSubAccount = append(s.TransactionOnSubAccount, newValue)
	return newValue
}
