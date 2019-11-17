package model

// Service provided by the ATM inside the session.
type ATMService18 struct {

	// Identification of the service variant.
	Identification *Max35Text `xml:"Id"`

	// Label of the service variant.
	Label *Max35Text `xml:"Labl,omitempty"`
}

func (a *ATMService18) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *ATMService18) SetLabel(value string) {
	a.Label = (*Max35Text)(&value)
}
