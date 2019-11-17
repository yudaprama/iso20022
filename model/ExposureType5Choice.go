package model

// Choice of format for the exposure information.
type ExposureType5Choice struct {

	// Collateral movement exposure type expressed as an ISO 20022 code.
	Code *ExposureType2Code `xml:"Cd"`

	// Collateral movement exposure type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (e *ExposureType5Choice) SetCode(value string) {
	e.Code = (*ExposureType2Code)(&value)
}

func (e *ExposureType5Choice) AddProprietary() *GenericIdentification20 {
	e.Proprietary = new(GenericIdentification20)
	return e.Proprietary
}
