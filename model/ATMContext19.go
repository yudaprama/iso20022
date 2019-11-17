package model

// Context in which the transfer is performed.
type ATMContext19 struct {

	// Unique identification of the customer session in which the transfer is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Fund transfer service requested by the ATM inside the session.
	Service *ATMService23 `xml:"Svc"`
}

func (a *ATMContext19) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext19) AddService() *ATMService23 {
	a.Service = new(ATMService23)
	return a.Service
}
