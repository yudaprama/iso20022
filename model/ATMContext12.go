package model

// Context in which the transaction is performed.
type ATMContext12 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Deposit service provided by the ATM inside the session.
	Service *ATMService13 `xml:"Svc,omitempty"`
}

func (a *ATMContext12) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext12) AddService() *ATMService13 {
	a.Service = new(ATMService13)
	return a.Service
}
