package model

// Choice between a standard code or proprietary code to specify the type of election movement.
type ElectionTypeFormat1Choice struct {

	// Standard code to specify the effect on the holdings of electing a corporate action option.
	Code *ElectionMovementType2Code `xml:"Cd"`

	// Proprietary identification of the type of election of a corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (e *ElectionTypeFormat1Choice) SetCode(value string) {
	e.Code = (*ElectionMovementType2Code)(&value)
}

func (e *ElectionTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	e.Proprietary = new(GenericIdentification20)
	return e.Proprietary
}
