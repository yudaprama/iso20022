package model

// Choice between a standard code or proprietary code to specify the type of offer.
type OfferTypeFormat3Choice struct {

	// Standard code to specify the conditions that apply to the offer.
	Code *OfferType2Code `xml:"Cd"`

	// Proprietary identification of the conditions that apply to the offer.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OfferTypeFormat3Choice) SetCode(value string) {
	o.Code = (*OfferType2Code)(&value)
}

func (o *OfferTypeFormat3Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
