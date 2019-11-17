package model

// Other accepted financial instrument's identification than ISIN.
type OtherIdentification1 struct {

	// Identification of a security.
	Identification *Max35Text `xml:"Id"`

	// Identifies the suffix of the security identification.
	Suffix *Max16Text `xml:"Sfx,omitempty"`

	// Type of the identification.
	Type *IdentificationSource3Choice `xml:"Tp"`
}

func (o *OtherIdentification1) SetIdentification(value string) {
	o.Identification = (*Max35Text)(&value)
}

func (o *OtherIdentification1) SetSuffix(value string) {
	o.Suffix = (*Max16Text)(&value)
}

func (o *OtherIdentification1) AddType() *IdentificationSource3Choice {
	o.Type = new(IdentificationSource3Choice)
	return o.Type
}
