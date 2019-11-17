package model

// Choice of format for the classification.
type ClassificationType32Choice struct {

	// ISO 10962 Classification of Financial Instrument (CFI)
	ClassificationFinancialInstrument *CFIOct2015Identifier `xml:"ClssfctnFinInstrm"`

	// Proprietary classification of financial instrument.
	AlternateClassification *GenericIdentification36 `xml:"AltrnClssfctn"`
}

func (c *ClassificationType32Choice) SetClassificationFinancialInstrument(value string) {
	c.ClassificationFinancialInstrument = (*CFIOct2015Identifier)(&value)
}

func (c *ClassificationType32Choice) AddAlternateClassification() *GenericIdentification36 {
	c.AlternateClassification = new(GenericIdentification36)
	return c.AlternateClassification
}
