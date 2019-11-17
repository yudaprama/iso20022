package model

// Choice of formats to  express the value of a rate.
type RateValueType6FormatChoice struct {

	// Standard code to specify the value of a rate.
	Code *RateValueType6Code `xml:"Cd"`

	// Proprietary code to  express the value of a rate.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RateValueType6FormatChoice) SetCode(value string) {
	r.Code = (*RateValueType6Code)(&value)
}

func (r *RateValueType6FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
