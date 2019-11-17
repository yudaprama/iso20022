package model

// Set of elements used to further detail the information related to the type of payment.
type MandateTypeInformation2 struct {

	// Agreement under which or rules under which the mandate resides.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`

	// Specifies the high level purpose of the mandate based on a set of pre-defined categories.
	CategoryPurpose *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`

	// Type of direct debit instruction.
	Classification *MandateClassification1Choice `xml:"Clssfctn,omitempty"`
}

func (m *MandateTypeInformation2) AddServiceLevel() *ServiceLevel8Choice {
	m.ServiceLevel = new(ServiceLevel8Choice)
	return m.ServiceLevel
}

func (m *MandateTypeInformation2) AddLocalInstrument() *LocalInstrument2Choice {
	m.LocalInstrument = new(LocalInstrument2Choice)
	return m.LocalInstrument
}

func (m *MandateTypeInformation2) AddCategoryPurpose() *CategoryPurpose1Choice {
	m.CategoryPurpose = new(CategoryPurpose1Choice)
	return m.CategoryPurpose
}

func (m *MandateTypeInformation2) AddClassification() *MandateClassification1Choice {
	m.Classification = new(MandateClassification1Choice)
	return m.Classification
}
