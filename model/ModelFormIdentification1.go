package model

// Identification of a model form.
type ModelFormIdentification1 struct {

	// Identification of the model form.
	Identification *ModelFormIdentification1Choice `xml:"Id"`

	// Version of the model form.
	Version *Max35Text `xml:"Vrsn,omitempty"`
}

func (m *ModelFormIdentification1) AddIdentification() *ModelFormIdentification1Choice {
	m.Identification = new(ModelFormIdentification1Choice)
	return m.Identification
}

func (m *ModelFormIdentification1) SetVersion(value string) {
	m.Version = (*Max35Text)(&value)
}
