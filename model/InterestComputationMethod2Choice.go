package model

// Choice between a code or a data source scheme to determine the interest computation method.
type InterestComputationMethod2Choice struct {

	// Code is used to determine the interest computation method.
	Code *InterestComputationMethod1Code `xml:"Cd"`

	// Interest computation method is determined using a data source scheme.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (i *InterestComputationMethod2Choice) SetCode(value string) {
	i.Code = (*InterestComputationMethod1Code)(&value)
}

func (i *InterestComputationMethod2Choice) AddProprietary() *GenericIdentification38 {
	i.Proprietary = new(GenericIdentification38)
	return i.Proprietary
}
