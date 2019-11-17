package model

// Account to or from which a securities entry is made.
type SubAccountIdentification36 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnSubAccount []*InvestmentFundTransactionsByFund3 `xml:"TxOnSubAcct,omitempty"`
}

func (s *SubAccountIdentification36) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification36) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification36) AddTransactionOnSubAccount() *InvestmentFundTransactionsByFund3 {
	newValue := new(InvestmentFundTransactionsByFund3)
	s.TransactionOnSubAccount = append(s.TransactionOnSubAccount, newValue)
	return newValue
}
