package model

// Choice of formats to  express the value of a rate.
type RateValueType2FormatChoice struct {

	// Standard code to specify the value of a rate.
	Code *RateValueType2Code `xml:"Cd"`

	// Proprietary code to  express the value of a rate.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RateValueType2FormatChoice) SetCode(value string) {
	r.Code = (*RateValueType2Code)(&value)
}

func (r *RateValueType2FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
