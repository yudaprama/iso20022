package model

// Specifies the details of the account and the role of the party.
type AccountAndParties2 struct {

	// Description of the account.
	Account *CustomerAccount1 `xml:"Acct"`

	// Specifies the role related to the account.
	Role []*AccountRole1 `xml:"Role"`

	// Additional information.
	AdditionalInformation []*Max256Text `xml:"AddtlInf,omitempty"`
}

func (a *AccountAndParties2) AddAccount() *CustomerAccount1 {
	a.Account = new(CustomerAccount1)
	return a.Account
}

func (a *AccountAndParties2) AddRole() *AccountRole1 {
	newValue := new(AccountRole1)
	a.Role = append(a.Role, newValue)
	return newValue
}

func (a *AccountAndParties2) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max256Text)(&value))
}
