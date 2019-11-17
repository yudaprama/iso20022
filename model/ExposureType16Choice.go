package model

// Choice of format for the exposure information.
type ExposureType16Choice struct {

	// Collateral movement exposure type expressed as an ISO 20022 code.
	Code *ExposureType4Code `xml:"Cd"`

	// Collateral movement exposure type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (e *ExposureType16Choice) SetCode(value string) {
	e.Code = (*ExposureType4Code)(&value)
}

func (e *ExposureType16Choice) AddProprietary() *GenericIdentification30 {
	e.Proprietary = new(GenericIdentification30)
	return e.Proprietary
}
