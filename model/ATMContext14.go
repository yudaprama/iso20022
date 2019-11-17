package model

// Context in which the inquiry is performed.
type ATMContext14 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService15 `xml:"Svc"`
}

func (a *ATMContext14) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext14) AddService() *ATMService15 {
	a.Service = new(ATMService15)
	return a.Service
}
