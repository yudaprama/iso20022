package model

// Choice between a code and or a data source scheme to determine the rate.
type RateType67Choice struct {

	// Rate expressed as an ISO 20022 code.
	Code *RateType1Code `xml:"Cd"`

	// Rate expressed as an a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType67Choice) SetCode(value string) {
	r.Code = (*RateType1Code)(&value)
}

func (r *RateType67Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
