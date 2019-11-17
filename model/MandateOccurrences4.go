package model

// Provides further details related to the duration of the mandate and the occurrence of the underlying transactions.
type MandateOccurrences4 struct {

	// Identifies the underlying transaction sequence as either recurring or one-off.
	SequenceType *SequenceType2Code `xml:"SeqTp"`

	// Regularity with which instructions are to be created and processed.
	Frequency *Frequency36Choice `xml:"Frqcy,omitempty"`

	// Length of time for which the mandate remains valid.
	Duration *DatePeriodDetails1 `xml:"Drtn,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`
}

func (m *MandateOccurrences4) SetSequenceType(value string) {
	m.SequenceType = (*SequenceType2Code)(&value)
}

func (m *MandateOccurrences4) AddFrequency() *Frequency36Choice {
	m.Frequency = new(Frequency36Choice)
	return m.Frequency
}

func (m *MandateOccurrences4) AddDuration() *DatePeriodDetails1 {
	m.Duration = new(DatePeriodDetails1)
	return m.Duration
}

func (m *MandateOccurrences4) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateOccurrences4) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}
