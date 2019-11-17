package model

// Account to or from which a securities entry is made.
type SubAccountIdentification5 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForSubAccount []*AggregateBalanceInformation4 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification5) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification5) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification5) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification5) AddBalanceForSubAccount() *AggregateBalanceInformation4 {
	newValue := new(AggregateBalanceInformation4)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}
