package model

// Context in which the inquiry is performed.
type ATMContext15 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService16 `xml:"Svc"`
}

func (a *ATMContext15) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext15) AddService() *ATMService16 {
	a.Service = new(ATMService16)
	return a.Service
}
