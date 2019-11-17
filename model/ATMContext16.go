package model

// Context in which the transaction is performed.
type ATMContext16 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Context in which the transaction is performed.
	Service *ATMService20 `xml:"Svc"`
}

func (a *ATMContext16) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext16) AddService() *ATMService20 {
	a.Service = new(ATMService20)
	return a.Service
}
