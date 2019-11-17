package model

// Account to or from which a securities entry is made.
type SubAccountIdentification1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForSubAccount []*AggregateBalanceInformation1 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification1) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification1) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification1) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification1) AddBalanceForSubAccount() *AggregateBalanceInformation1 {
	newValue := new(AggregateBalanceInformation1)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
