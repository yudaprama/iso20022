package model

// Set of elements that further identifies the type of local instruments being requested by the initiating party.
type LocalInstrument1Choice struct {

	// Specifies the local instrument published in an external local instrument code list.
	Code *ExternalLocalInstrumentCode `xml:"Cd"`

	// Specifies the local instrument as a proprietary code.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (l *LocalInstrument1Choice) SetCode(value string) {
	l.Code = (*ExternalLocalInstrumentCode)(&value)
}

func (l *LocalInstrument1Choice) SetProprietary(value string) {
	l.Proprietary = (*Max35Text)(&value)
}
