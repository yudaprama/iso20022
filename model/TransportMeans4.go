package model

// Describes the multimodal or the individual transport of goods.
type TransportMeans4 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport6 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans4) AddIndividualTransport() *SingleTransport6 {
	t.IndividualTransport = new(SingleTransport6)
	return t.IndividualTransport
}

func (t *TransportMeans4) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}
