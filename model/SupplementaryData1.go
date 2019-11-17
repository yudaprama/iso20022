package model

// Additional information that can not be captured in the structured fields and/or any other specific block.
type SupplementaryData1 struct {

	// Unambiguous reference to the location where the supplementary data must be inserted in the message instance.
	// In the case of XML, this is expressed by a valid XPath.
	PlaceAndName *Max350Text `xml:"PlcAndNm,omitempty"`

	// Technical element wrapping the supplementary data.
	Envelope *SupplementaryDataEnvelope1 `xml:"Envlp"`
}

func (s *SupplementaryData1) SetPlaceAndName(value string) {
	s.PlaceAndName = (*Max350Text)(&value)
}

func (s *SupplementaryData1) AddEnvelope() *SupplementaryDataEnvelope1 {
	s.Envelope = new(SupplementaryDataEnvelope1)
	return s.Envelope
}
