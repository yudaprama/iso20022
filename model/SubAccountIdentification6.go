package model

// Account to or from which a securities entry is made.
type SubAccountIdentification6 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnSubAccount []*InvestmentFundTransactionsByFund2 `xml:"TxOnSubAcct,omitempty"`
}

func (s *SubAccountIdentification6) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification6) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification6) AddTransactionOnSubAccount() *InvestmentFundTransactionsByFund2 {
	newValue := new(InvestmentFundTransactionsByFund2)
	s.TransactionOnSubAccount = append(s.TransactionOnSubAccount, newValue)
	return newValue
}
