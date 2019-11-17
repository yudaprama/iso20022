package model

// Choice of formats to  express the effect on the holdings of electing a Corporate Action option.
type ElectionMovementType1FormatChoice struct {

	// Standard code to  specify the effect on the holdings of electing a Corporate Action option.
	Code *ElectionMovementType1Code `xml:"Cd"`

	// Proprietary code to  express the effect on the holdings of electing a Corporate Action option.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (e *ElectionMovementType1FormatChoice) SetCode(value string) {
	e.Code = (*ElectionMovementType1Code)(&value)
}

func (e *ElectionMovementType1FormatChoice) AddProprietary() *GenericIdentification13 {
	e.Proprietary = new(GenericIdentification13)
	return e.Proprietary
}
