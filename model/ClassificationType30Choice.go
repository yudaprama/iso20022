package model

// Choice of format for the classification.
type ClassificationType30Choice struct {

	// ISO 10962 Classification of Financial Instrument (CFI)
	ClassificationFinancialInstrument *CFIIdentifier `xml:"ClssfctnFinInstrm"`

	// Proprietary classification of financial instrument.
	AlternateClassification *GenericIdentification36 `xml:"AltrnClssfctn"`
}

func (c *ClassificationType30Choice) SetClassificationFinancialInstrument(value string) {
	c.ClassificationFinancialInstrument = (*CFIIdentifier)(&value)
}

func (c *ClassificationType30Choice) AddAlternateClassification() *GenericIdentification36 {
	c.AlternateClassification = new(GenericIdentification36)
	return c.AlternateClassification
}
