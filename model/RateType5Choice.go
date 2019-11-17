package model

// Choice between a code and or a data source scheme to determine the rate.
type RateType5Choice struct {

	// Rate expressed as an ISO 20022 code.
	Code *RateType1Code `xml:"Cd"`

	// Rate expressed as an a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType5Choice) SetCode(value string) {
	r.Code = (*RateType1Code)(&value)
}

func (r *RateType5Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
