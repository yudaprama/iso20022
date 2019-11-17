package model

// Choice of format for the exposure information.
type ExposureType9Choice struct {

	// Collateral movement exposure type expressed as an ISO 20022 code.
	Code *ExposureType3Code `xml:"Cd"`

	// Collateral movement exposure type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (e *ExposureType9Choice) SetCode(value string) {
	e.Code = (*ExposureType3Code)(&value)
}

func (e *ExposureType9Choice) AddProprietary() *GenericIdentification38 {
	e.Proprietary = new(GenericIdentification38)
	return e.Proprietary
}
