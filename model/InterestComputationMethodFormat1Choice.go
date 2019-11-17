package model

// Choice between a standard code or proprietary code to specify the type of interest computation method.
type InterestComputationMethodFormat1Choice struct {

	// Standard code to specify the method used to compute accruing interest of a financial instrument.
	Code *InterestComputationMethod2Code `xml:"Cd"`

	// Proprietary identification of the format of interest computation method.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *InterestComputationMethodFormat1Choice) SetCode(value string) {
	i.Code = (*InterestComputationMethod2Code)(&value)
}

func (i *InterestComputationMethodFormat1Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
