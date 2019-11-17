package model

// Institution in charge of managing the ATM.
type Acquirer8 struct {

	// Identification of the ATM manager.
	Identification *Max35Text `xml:"Id"`

	// Software version of the application.
	ApplicationVersion *Max35Text `xml:"ApplVrsn,omitempty"`
}

func (a *Acquirer8) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Acquirer8) SetApplicationVersion(value string) {
	a.ApplicationVersion = (*Max35Text)(&value)
}
