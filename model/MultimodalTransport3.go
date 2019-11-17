package model

// Information related to multimodal transportation of goods.
type MultimodalTransport3 struct {

	// Identifies the location where the goods are take in charge for transportation.
	TakingInCharge *Max35Text `xml:"TakngInChrg"`

	// Identifies the location of the final destination of the goods.
	PlaceOfFinalDestination *Max35Text `xml:"PlcOfFnlDstn"`
}

func (m *MultimodalTransport3) SetTakingInCharge(value string) {
	m.TakingInCharge = (*Max35Text)(&value)
}

func (m *MultimodalTransport3) SetPlaceOfFinalDestination(value string) {
	m.PlaceOfFinalDestination = (*Max35Text)(&value)
}
