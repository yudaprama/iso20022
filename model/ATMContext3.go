package model

// Context in which the transaction is performed.
type ATMContext3 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Service provided by the ATM inside the session.
	Service *ATMService3 `xml:"Svc"`
}

func (a *ATMContext3) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext3) AddService() *ATMService3 {
	a.Service = new(ATMService3)
	return a.Service
}
