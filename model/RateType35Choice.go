package model

// Choice between a code and or a data source scheme to determine the rate.
type RateType35Choice struct {

	// Rate expressed as an ISO 20022 code.
	Code *RateType1Code `xml:"Cd"`

	// Rate expressed as an a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType35Choice) SetCode(value string) {
	r.Code = (*RateType1Code)(&value)
}

func (r *RateType35Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
