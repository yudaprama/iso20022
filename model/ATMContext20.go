package model

// Context of the exception.
type ATMContext20 struct {

	// Unique identification of the customer session in which the exception occurred.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Service provided by the ATM inside the session.
	Service *ATMService24 `xml:"Svc"`
}

func (a *ATMContext20) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext20) AddService() *ATMService24 {
	a.Service = new(ATMService24)
	return a.Service
}
