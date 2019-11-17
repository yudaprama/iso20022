package model

// Customer involved in a withdrawal transaction.
type ATMCustomer1 struct {

	// Profile of the customer selected to perform the withdrawal.
	Profile *ATMCustomerProfile1 `xml:"Prfl,omitempty"`

	// Language selected by the customer at the ATM for the customer session. Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	SelectedLanguage *LanguageCode `xml:"SelctdLang,omitempty"`

	// Method and data intended to be used for this transaction to authenticate the customer and its card.
	Authentication []*CardholderAuthentication8 `xml:"Authntcn"`

	// Result of the customer authentication for this transaction.
	AuthenticationResult []*TransactionVerificationResult5 `xml:"AuthntcnRslt,omitempty"`
}

func (a *ATMCustomer1) AddProfile() *ATMCustomerProfile1 {
	a.Profile = new(ATMCustomerProfile1)
	return a.Profile
}

func (a *ATMCustomer1) SetSelectedLanguage(value string) {
	a.SelectedLanguage = (*LanguageCode)(&value)
}

func (a *ATMCustomer1) AddAuthentication() *CardholderAuthentication8 {
	newValue := new(CardholderAuthentication8)
	a.Authentication = append(a.Authentication, newValue)
	return newValue
}

func (a *ATMCustomer1) AddAuthenticationResult() *TransactionVerificationResult5 {
	newValue := new(TransactionVerificationResult5)
	a.AuthenticationResult = append(a.AuthenticationResult, newValue)
	return newValue
}
