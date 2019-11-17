package model

// Specifies the applicable Incoterm and associated location.
type Incoterms4 struct {

	// Specifies the Incoterms.
	IncotermsCode *Incoterms4Choice `xml:"IncotrmsCd"`

	// Location where the Incoterms are actioned. Reference UN/ECE Recommendation 16 - LOCODE - Code for Trade and Transport Locations (www.unece.org/cefact/recommendations).
	Location *Max70Text `xml:"Lctn,omitempty"`
}

func (i *Incoterms4) AddIncotermsCode() *Incoterms4Choice {
	i.IncotermsCode = new(Incoterms4Choice)
	return i.IncotermsCode
}

func (i *Incoterms4) SetLocation(value string) {
	i.Location = (*Max70Text)(&value)
}
