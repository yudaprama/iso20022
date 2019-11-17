package model

// Choice between a code or a data source scheme to determine the business process type.
type BusinessProcessType1Choice struct {

	// Business process type is identified using a code.
	Code *BusinessProcessType1Code `xml:"Cd"`

	// Type of business process expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (b *BusinessProcessType1Choice) SetCode(value string) {
	b.Code = (*BusinessProcessType1Code)(&value)
}

func (b *BusinessProcessType1Choice) AddProprietary() *GenericIdentification38 {
	b.Proprietary = new(GenericIdentification38)
	return b.Proprietary
}
