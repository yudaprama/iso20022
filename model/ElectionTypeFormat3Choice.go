package model

// Choice between a standard code or proprietary code to specify the type of election movement.
type ElectionTypeFormat3Choice struct {

	// Standard code to specify the effect on the holdings of electing a corporate action option.
	Code *ElectionMovementType2Code `xml:"Cd"`

	// Proprietary identification of the type of election of a corporate action option.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (e *ElectionTypeFormat3Choice) SetCode(value string) {
	e.Code = (*ElectionMovementType2Code)(&value)
}

func (e *ElectionTypeFormat3Choice) AddProprietary() *GenericIdentification30 {
	e.Proprietary = new(GenericIdentification30)
	return e.Proprietary
}
