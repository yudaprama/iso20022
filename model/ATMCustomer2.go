package model

// Customer involved in a withdrawal transaction.
type ATMCustomer2 struct {

	// Profile of the customer selected to perform the withdrawal.
	Profile *ATMCustomerProfile2 `xml:"Prfl,omitempty"`

	// Result of the customer authentication for this transaction.
	AuthenticationResult []*TransactionVerificationResult5 `xml:"AuthntcnRslt,omitempty"`
}

func (a *ATMCustomer2) AddProfile() *ATMCustomerProfile2 {
	a.Profile = new(ATMCustomerProfile2)
	return a.Profile
}

func (a *ATMCustomer2) AddAuthenticationResult() *TransactionVerificationResult5 {
	newValue := new(TransactionVerificationResult5)
	a.AuthenticationResult = append(a.AuthenticationResult, newValue)
	return newValue
}
