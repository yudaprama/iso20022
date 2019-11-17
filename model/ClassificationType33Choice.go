package model

// Choice of format for the classification.
type ClassificationType33Choice struct {

	// ISO 10962 Classification of Financial Instrument (CFI)
	ClassificationFinancialInstrument *CFIOct2015Identifier `xml:"ClssfctnFinInstrm"`

	// Proprietary classification of financial instrument.
	AlternateClassification *GenericIdentification86 `xml:"AltrnClssfctn"`
}

func (c *ClassificationType33Choice) SetClassificationFinancialInstrument(value string) {
	c.ClassificationFinancialInstrument = (*CFIOct2015Identifier)(&value)
}

func (c *ClassificationType33Choice) AddAlternateClassification() *GenericIdentification86 {
	c.AlternateClassification = new(GenericIdentification86)
	return c.AlternateClassification
}
