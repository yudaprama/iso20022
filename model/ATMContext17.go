package model

// Context in which the transaction is performed.
type ATMContext17 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService21 `xml:"Svc"`
}

func (a *ATMContext17) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext17) AddService() *ATMService21 {
	a.Service = new(ATMService21)
	return a.Service
}
