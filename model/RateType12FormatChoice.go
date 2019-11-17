package model

// Choice of formats to  express the type of rate.
type RateType12FormatChoice struct {

	// Standard code to specify the type of rate.
	Code *RateType12Code `xml:"Cd"`

	// Proprietary code to  express the type of rate.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RateType12FormatChoice) SetCode(value string) {
	r.Code = (*RateType12Code)(&value)
}

func (r *RateType12FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
