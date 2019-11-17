package model

// Choice of format for the classification.
type ClassificationType1Choice struct {

	// ISO 10962 Classification of Financial Instrument (CFI)
	ClassificationFinancialInstrument *CFIOct2015Identifier `xml:"ClssfctnFinInstrm"`

	// Proprietary classification of financial instrument.
	AlternateClassification *GenericIdentification1 `xml:"AltrnClssfctn"`
}

func (c *ClassificationType1Choice) SetClassificationFinancialInstrument(value string) {
	c.ClassificationFinancialInstrument = (*CFIOct2015Identifier)(&value)
}

func (c *ClassificationType1Choice) AddAlternateClassification() *GenericIdentification1 {
	c.AlternateClassification = new(GenericIdentification1)
	return c.AlternateClassification
}
