package model

// Choice between a standard code or proprietary code to specify the type of offer.
type OfferTypeFormat11Choice struct {

	// Standard code to specify the conditions that apply to the offer.
	Code *OfferType3Code `xml:"Cd"`

	// Proprietary identification of the conditions that apply to the offer.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OfferTypeFormat11Choice) SetCode(value string) {
	o.Code = (*OfferType3Code)(&value)
}

func (o *OfferTypeFormat11Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
