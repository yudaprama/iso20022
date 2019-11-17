package model

// Choice between a standard code or proprietary code to specify the type of the additional business process.
type AdditionalBusinessProcessFormat1Choice struct {

	// Standard code to specify the additional business process linked to a corporate action event.
	Code *AdditionalBusinessProcess1Code `xml:"Cd"`

	// Proprietary identification of the additional business process linked to a corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat1Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess1Code)(&value)
}

func (a *AdditionalBusinessProcessFormat1Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
