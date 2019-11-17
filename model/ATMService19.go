package model

// Service allowed on the account.
type ATMService19 struct {

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType8Code `xml:"SvcTp"`

	// Variant of the service.
	ServiceVariant []*ATMService18 `xml:"SvcVarnt,omitempty"`

	// Limits of amounts.
	Limits []*ATMTransactionAmounts6 `xml:"Lmts,omitempty"`
}

func (a *ATMService19) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType8Code)(&value)
}

func (a *ATMService19) AddServiceVariant() *ATMService18 {
	newValue := new(ATMService18)
	a.ServiceVariant = append(a.ServiceVariant, newValue)
	return newValue
}

func (a *ATMService19) AddLimits() *ATMTransactionAmounts6 {
	newValue := new(ATMTransactionAmounts6)
	a.Limits = append(a.Limits, newValue)
	return newValue
}
