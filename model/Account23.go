package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account23 struct {

	// Identification of the account.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Information about the account to which the existing account is to be linked.
	RelatedAccountDetails *GenericIdentification1 `xml:"RltdAcctDtls,omitempty"`
}

func (a *Account23) SetAccountIdentification(value string) {
	a.AccountIdentification = (*Max35Text)(&value)
}

func (a *Account23) AddRelatedAccountDetails() *GenericIdentification1 {
	a.RelatedAccountDetails = new(GenericIdentification1)
	return a.RelatedAccountDetails
}
