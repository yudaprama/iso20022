package model

// Context in which the transaction is performed.
type ATMContext9 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService10 `xml:"Svc,omitempty"`
}

func (a *ATMContext9) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext9) AddService() *ATMService10 {
	a.Service = new(ATMService10)
	return a.Service
}
