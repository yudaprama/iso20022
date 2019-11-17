package model

// Context in which the transfer is performed.
type ATMContext18 struct {

	// Unique identification of the customer session in which the transfer is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Fund transfer service requested by the ATM inside the session.
	Service *ATMService22 `xml:"Svc"`
}

func (a *ATMContext18) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext18) AddService() *ATMService22 {
	a.Service = new(ATMService22)
	return a.Service
}
