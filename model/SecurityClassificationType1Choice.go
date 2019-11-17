package model

// Choice between a CFI code or an other type of identification for the classification of the financial instrument.
type SecurityClassificationType1Choice struct {

	// Classification type of the financial instrument, as per the ISO 10962 Classification of Financial Instrument (CFI) codification.
	CFI *CFIIdentifier `xml:"CFI"`

	// Other type of classification of the financial instrument.
	AlternateClassification *GenericIdentification3 `xml:"AltrnClssfctn"`
}

func (s *SecurityClassificationType1Choice) SetCFI(value string) {
	s.CFI = (*CFIIdentifier)(&value)
}

func (s *SecurityClassificationType1Choice) AddAlternateClassification() *GenericIdentification3 {
	s.AlternateClassification = new(GenericIdentification3)
	return s.AlternateClassification
}
