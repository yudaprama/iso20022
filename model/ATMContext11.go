package model

// Context in which the transaction is performed.
type ATMContext11 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Deposit service provided by the ATM inside the session.
	Service *ATMService12 `xml:"Svc,omitempty"`
}

func (a *ATMContext11) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext11) AddService() *ATMService12 {
	a.Service = new(ATMService12)
	return a.Service
}
