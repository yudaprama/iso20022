package model

// Describes the multimodal or the individual transport of goods.
type TransportMeans3 struct {

	// Code specifying the transport mode for the delivery of the consignment, such as by air, sea, rail, road or inland waterway. Reference UN/ECE Recommendation 19 - Code for Modes of Transport (www.unece.org/cefact/recommendations).
	ModeCode *Max4Text `xml:"MdCd,omitempty"`

	// Unique identification of the means of transport, such as the International Maritime Organization number of a vessel.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Name, expressed as text, of the means of transport.
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (t *TransportMeans3) SetModeCode(value string) {
	t.ModeCode = (*Max4Text)(&value)
}

func (t *TransportMeans3) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}

func (t *TransportMeans3) SetName(value string) {
	t.Name = (*Max35Text)(&value)
}
