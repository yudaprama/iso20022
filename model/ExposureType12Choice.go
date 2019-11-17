package model

// Choice of format for the exposure information.
type ExposureType12Choice struct {

	// Collateral movement exposure type expressed as an ISO 20022 code.
	Code *ExposureType4Code `xml:"Cd"`

	// Collateral movement exposure type expressed as a proprietary code.
	Proprietary *GenericIdentification40 `xml:"Prtry"`
}

func (e *ExposureType12Choice) SetCode(value string) {
	e.Code = (*ExposureType4Code)(&value)
}

func (e *ExposureType12Choice) AddProprietary() *GenericIdentification40 {
	e.Proprietary = new(GenericIdentification40)
	return e.Proprietary
}
