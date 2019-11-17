package model

// Services allowed for the customer's profile.
type ATMService17 struct {

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType8Code `xml:"SvcTp"`

	// Variant of the service.
	ServiceVariant []*ATMService18 `xml:"SvcVarnt,omitempty"`

	// Limits of amounts.
	Limits []*ATMTransactionAmounts6 `xml:"Lmts,omitempty"`

	// Preferred withdrawal transaction chosen by the customer.
	PreferredWithdrawal *ATMTransaction8 `xml:"PrefrdWdrwl,omitempty"`
}

func (a *ATMService17) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType8Code)(&value)
}

func (a *ATMService17) AddServiceVariant() *ATMService18 {
	newValue := new(ATMService18)
	a.ServiceVariant = append(a.ServiceVariant, newValue)
	return newValue
}

func (a *ATMService17) AddLimits() *ATMTransactionAmounts6 {
	newValue := new(ATMTransactionAmounts6)
	a.Limits = append(a.Limits, newValue)
	return newValue
}

func (a *ATMService17) AddPreferredWithdrawal() *ATMTransaction8 {
	a.PreferredWithdrawal = new(ATMTransaction8)
	return a.PreferredWithdrawal
}
