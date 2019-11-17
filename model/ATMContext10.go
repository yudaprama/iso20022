package model

// Context in which the transaction is performed.
type ATMContext10 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Deposit service provided by the ATM inside the session.
	Service *ATMService11 `xml:"Svc,omitempty"`
}

func (a *ATMContext10) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext10) AddService() *ATMService11 {
	a.Service = new(ATMService11)
	return a.Service
}
