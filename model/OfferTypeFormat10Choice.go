package model

// Choice between a standard code or proprietary code to specify the type of offer.
type OfferTypeFormat10Choice struct {

	// Standard code to specify the conditions that apply to the offer.
	Code *OfferType3Code `xml:"Cd"`

	// Proprietary identification of the conditions that apply to the offer.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OfferTypeFormat10Choice) SetCode(value string) {
	o.Code = (*OfferType3Code)(&value)
}

func (o *OfferTypeFormat10Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}
