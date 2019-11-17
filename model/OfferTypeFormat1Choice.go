package model

// Choice between a standard code or proprietary code to specify the type of offer.
type OfferTypeFormat1Choice struct {

	// Standard code to specify the conditions that apply to the offer.
	Code *OfferType1Code `xml:"Cd"`

	// Proprietary identification of the conditions that apply to the offer.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OfferTypeFormat1Choice) SetCode(value string) {
	o.Code = (*OfferType1Code)(&value)
}

func (o *OfferTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
