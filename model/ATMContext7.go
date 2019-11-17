package model

// Context in which the transaction is performed.
type ATMContext7 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Context in which the transaction is performed.
	Service *ATMService8 `xml:"Svc"`
}

func (a *ATMContext7) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext7) AddService() *ATMService8 {
	a.Service = new(ATMService8)
	return a.Service
}
