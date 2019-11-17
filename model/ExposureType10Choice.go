package model

// Choice of format for the exposure information.
type ExposureType10Choice struct {

	// Collateral movement exposure type expressed as an ISO 20022 code.
	Code *ExposureType4Code `xml:"Cd"`

	// Collateral movement exposure type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (e *ExposureType10Choice) SetCode(value string) {
	e.Code = (*ExposureType4Code)(&value)
}

func (e *ExposureType10Choice) AddProprietary() *GenericIdentification20 {
	e.Proprietary = new(GenericIdentification20)
	return e.Proprietary
}
