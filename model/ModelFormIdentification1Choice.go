package model

// Choice of format for the model form identification.
type ModelFormIdentification1Choice struct {

	// Model form identification.
	//
	Code *ExternalModelFormIdentification1Code `xml:"Cd"`

	// Model form identification expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (m *ModelFormIdentification1Choice) SetCode(value string) {
	m.Code = (*ExternalModelFormIdentification1Code)(&value)
}

func (m *ModelFormIdentification1Choice) AddProprietary() *GenericIdentification1 {
	m.Proprietary = new(GenericIdentification1)
	return m.Proprietary
}
