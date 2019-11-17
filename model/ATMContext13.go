package model

// Context in which the transaction is performed.
type ATMContext13 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Service provided by the ATM inside the session.
	Service *ATMService14 `xml:"Svc"`
}

func (a *ATMContext13) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext13) AddService() *ATMService14 {
	a.Service = new(ATMService14)
	return a.Service
}
