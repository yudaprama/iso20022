package model

// Choice between a response type code and a proprietary code.
type ResponseType1Choice struct {

	// Provides a margin call response type using an ISO 20022 code.
	Code *MarginCallResponse1Code `xml:"Cd"`

	// Provides a margin call response type using a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *ResponseType1Choice) SetCode(value string) {
	r.Code = (*MarginCallResponse1Code)(&value)
}

func (r *ResponseType1Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
