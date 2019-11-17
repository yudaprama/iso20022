package model

// Identification of the entity issuing the command.
type ATMCommandIdentification1 struct {

	// Identification of the entity issuing the command.
	Origin *Max35Text `xml:"Orgn,omitempty"`

	// Unique identification of the command for the issuer of the command.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Identification or address of the host performing the command.
	Processor *Max140Text `xml:"Prcr,omitempty"`
}

func (a *ATMCommandIdentification1) SetOrigin(value string) {
	a.Origin = (*Max35Text)(&value)
}

func (a *ATMCommandIdentification1) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *ATMCommandIdentification1) SetProcessor(value string) {
	a.Processor = (*Max140Text)(&value)
}
