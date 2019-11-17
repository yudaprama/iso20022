package model

// Specifies the applicable Incoterm and associated location.
type Incoterms1 struct {

	// Specifies the applicable Incoterm by means of a code.
	Code *Incoterms1Code `xml:"Cd"`

	// Specifies Incoterm not present in code list.
	Other *Max35Text `xml:"Othr"`

	// Location where the Incoterms are actioned.
	Location *Max35Text `xml:"Lctn,omitempty"`
}

func (i *Incoterms1) SetCode(value string) {
	i.Code = (*Incoterms1Code)(&value)
}

func (i *Incoterms1) SetOther(value string) {
	i.Other = (*Max35Text)(&value)
}

func (i *Incoterms1) SetLocation(value string) {
	i.Location = (*Max35Text)(&value)
}
