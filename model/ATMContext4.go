package model

// Context in which the transaction is performed.
type ATMContext4 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService4 `xml:"Svc"`
}

func (a *ATMContext4) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext4) AddService() *ATMService4 {
	a.Service = new(ATMService4)
	return a.Service
}
