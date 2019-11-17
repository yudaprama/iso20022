package model

// Context in which the inquiry is performed.
type ATMContext5 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService5 `xml:"Svc"`
}

func (a *ATMContext5) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext5) AddService() *ATMService5 {
	a.Service = new(ATMService5)
	return a.Service
}
