package model

// Set of elements that further identifies the type of local instruments being requested by the initiating party.
type LocalInstrument2Choice struct {

	// Specifies the local instrument, as published in an external local instrument code list.
	Code *ExternalLocalInstrument1Code `xml:"Cd"`

	// Specifies the local instrument, as a proprietary code.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (l *LocalInstrument2Choice) SetCode(value string) {
	l.Code = (*ExternalLocalInstrument1Code)(&value)
}

func (l *LocalInstrument2Choice) SetProprietary(value string) {
	l.Proprietary = (*Max35Text)(&value)
}
