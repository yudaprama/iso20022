package model

// Choice between a standard code or proprietary code to specify the type of offer.
type OfferTypeFormat5Choice struct {

	// Standard code to specify the conditions that apply to the offer.
	Code *OfferType3Code `xml:"Cd"`

	// Proprietary identification of the conditions that apply to the offer.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OfferTypeFormat5Choice) SetCode(value string) {
	o.Code = (*OfferType3Code)(&value)
}

func (o *OfferTypeFormat5Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
