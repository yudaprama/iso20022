package model

// Set of elements used to further detail the information related to the type of payment.
type MandateTypeInformation1 struct {

	// Agreement under which or rules under which the mandate resides.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
}

func (m *MandateTypeInformation1) AddServiceLevel() *ServiceLevel8Choice {
	m.ServiceLevel = new(ServiceLevel8Choice)
	return m.ServiceLevel
}

func (m *MandateTypeInformation1) AddLocalInstrument() *LocalInstrument2Choice {
	m.LocalInstrument = new(LocalInstrument2Choice)
	return m.LocalInstrument
}
