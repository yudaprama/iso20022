package model

// Other accepted financial instrument's identification than ISIN.
type OtherIdentification3 struct {

	// Identification of a security.
	Identification *RestrictedFINXMax31Text `xml:"Id"`

	// Identifies the suffix of the security identification.
	Suffix *Max16Text `xml:"Sfx,omitempty"`

	// Type of the identification.
	Type *IdentificationSource3Choice `xml:"Tp"`
}

func (o *OtherIdentification3) SetIdentification(value string) {
	o.Identification = (*RestrictedFINXMax31Text)(&value)
}

func (o *OtherIdentification3) SetSuffix(value string) {
	o.Suffix = (*Max16Text)(&value)
}

func (o *OtherIdentification3) AddType() *IdentificationSource3Choice {
	o.Type = new(IdentificationSource3Choice)
	return o.Type
}
