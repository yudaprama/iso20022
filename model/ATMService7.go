package model

// Services allowed for the customer's profile.
type ATMService7 struct {

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType4Code `xml:"SvcTp"`

	// Limits of amounts.
	Limits []*ATMTransactionAmounts3 `xml:"Lmts,omitempty"`

	// Preferred withdrawal transaction chosen by the customer.
	PreferredWithdrawal *ATMTransaction8 `xml:"PrefrdWdrwl,omitempty"`
}

func (a *ATMService7) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType4Code)(&value)
}

func (a *ATMService7) AddLimits() *ATMTransactionAmounts3 {
	newValue := new(ATMTransactionAmounts3)
	a.Limits = append(a.Limits, newValue)
	return newValue
}

func (a *ATMService7) AddPreferredWithdrawal() *ATMTransaction8 {
	a.PreferredWithdrawal = new(ATMTransaction8)
	return a.PreferredWithdrawal
}
