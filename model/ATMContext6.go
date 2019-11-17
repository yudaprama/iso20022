package model

// Context in which the inquiry is performed.
type ATMContext6 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService6 `xml:"Svc"`
}

func (a *ATMContext6) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext6) AddService() *ATMService6 {
	a.Service = new(ATMService6)
	return a.Service
}
