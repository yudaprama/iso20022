package model

// Choice of format for the classification.
type ClassificationType2Choice struct {

	// ISO 10962 Classification of Financial Instrument (CFI)
	ClassificationFinancialInstrument *CFIIdentifier `xml:"ClssfctnFinInstrm"`

	// Proprietary classification of financial instrument.
	AlternateClassification *GenericIdentification19 `xml:"AltrnClssfctn"`
}

func (c *ClassificationType2Choice) SetClassificationFinancialInstrument(value string) {
	c.ClassificationFinancialInstrument = (*CFIIdentifier)(&value)
}

func (c *ClassificationType2Choice) AddAlternateClassification() *GenericIdentification19 {
	c.AlternateClassification = new(GenericIdentification19)
	return c.AlternateClassification
}
