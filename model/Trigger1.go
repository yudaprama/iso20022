package model

// Trigger parameters.
type Trigger1 struct {

	// Details related to the date on which a variation is effective.
	DateChoice *FixedOrRecurrentDate1Choice `xml:"DtChc,omitempty"`

	// Details related to the documentary event on which a variation is triggered.
	DocumentaryEvent []*Document10 `xml:"DcmntryEvt,omitempty"`
}

func (t *Trigger1) AddDateChoice() *FixedOrRecurrentDate1Choice {
	t.DateChoice = new(FixedOrRecurrentDate1Choice)
	return t.DateChoice
}

func (t *Trigger1) AddDocumentaryEvent() *Document10 {
	newValue := new(Document10)
	t.DocumentaryEvent = append(t.DocumentaryEvent, newValue)
	return newValue
}
