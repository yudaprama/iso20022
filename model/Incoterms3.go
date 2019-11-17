package model

// Specifies the applicable Incoterm and associated location.
type Incoterms3 struct {

	// Specifies the Incoterms.
	IncotermsCode *Incoterms4Choice `xml:"IncotrmsCd"`

	// Location where the Incoterms are actioned. Reference UN/ECE Recommendation 16 - LOCODE - Code for Trade and Transport Locations (www.unece.org/cefact/recommendations).
	Location *Max35Text `xml:"Lctn,omitempty"`
}

func (i *Incoterms3) AddIncotermsCode() *Incoterms4Choice {
	i.IncotermsCode = new(Incoterms4Choice)
	return i.IncotermsCode
}

func (i *Incoterms3) SetLocation(value string) {
	i.Location = (*Max35Text)(&value)
}
