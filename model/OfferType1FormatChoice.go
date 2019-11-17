package model

// Choice of formats to  express the conditions that apply to the offer.
type OfferType1FormatChoice struct {

	// Standard code to  specify the conditions that apply to the offer.
	Code *OfferType1Code `xml:"Cd"`

	// Proprietary code to  express the conditions that apply to the offer.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (o *OfferType1FormatChoice) SetCode(value string) {
	o.Code = (*OfferType1Code)(&value)
}

func (o *OfferType1FormatChoice) AddProprietary() *GenericIdentification13 {
	o.Proprietary = new(GenericIdentification13)
	return o.Proprietary
}
