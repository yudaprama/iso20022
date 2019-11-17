package model

// Choice between a code and or a data source scheme to determine the rate.
type RateType19Choice struct {

	// Rate is defined using a code.
	Code *RateType1Code `xml:"Cd"`

	// Rate is determined using a data source scheme.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *RateType19Choice) SetCode(value string) {
	r.Code = (*RateType1Code)(&value)
}

func (r *RateType19Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
