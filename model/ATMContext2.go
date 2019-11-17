package model

// Context in which the transaction is performed.
type ATMContext2 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService2 `xml:"Svc,omitempty"`
}

func (a *ATMContext2) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext2) AddService() *ATMService2 {
	a.Service = new(ATMService2)
	return a.Service
}
