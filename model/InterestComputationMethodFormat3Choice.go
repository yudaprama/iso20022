package model

// Choice between a standard code or proprietary code to specify the type of interest computation method.
type InterestComputationMethodFormat3Choice struct {

	// Standard code to specify the method used to compute accruing interest of a financial instrument.
	Code *InterestComputationMethod2Code `xml:"Cd"`

	// Proprietary identification of the format of interest computation method.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (i *InterestComputationMethodFormat3Choice) SetCode(value string) {
	i.Code = (*InterestComputationMethod2Code)(&value)
}

func (i *InterestComputationMethodFormat3Choice) AddProprietary() *GenericIdentification38 {
	i.Proprietary = new(GenericIdentification38)
	return i.Proprietary
}
