package model

// Specifies the applicable Incoterm and associated location.
type Incoterms2 struct {

	// Specifies the applicable Incoterm by means of a code.
	Code *Incoterms1Code `xml:"Cd"`

	// Specifies Incoterm not present in code list.
	Other *Max35Text `xml:"Othr"`

	// Location where the Incoterms are actioned.
	Location *Max35Text `xml:"Lctn"`
}

func (i *Incoterms2) SetCode(value string) {
	i.Code = (*Incoterms1Code)(&value)
}

func (i *Incoterms2) SetOther(value string) {
	i.Other = (*Max35Text)(&value)
}

func (i *Incoterms2) SetLocation(value string) {
	i.Location = (*Max35Text)(&value)
}
