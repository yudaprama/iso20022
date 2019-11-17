package model

// Choice of format for the exposure information.
type ExposureType17Choice struct {

	// Collateral movement exposure type expressed as an ISO 20022 code.
	Code *ExposureType4Code `xml:"Cd"`

	// Collateral movement exposure type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (e *ExposureType17Choice) SetCode(value string) {
	e.Code = (*ExposureType4Code)(&value)
}

func (e *ExposureType17Choice) AddProprietary() *GenericIdentification47 {
	e.Proprietary = new(GenericIdentification47)
	return e.Proprietary
}
