package model

// Context in which the transaction is performed.
type ATMContext8 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService9 `xml:"Svc,omitempty"`
}

func (a *ATMContext8) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext8) AddService() *ATMService9 {
	a.Service = new(ATMService9)
	return a.Service
}
