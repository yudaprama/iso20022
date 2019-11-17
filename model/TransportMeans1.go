package model

// Describes the multimodal or the individual transport of goods.
type TransportMeans1 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport4 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans1) AddIndividualTransport() *SingleTransport4 {
	t.IndividualTransport = new(SingleTransport4)
	return t.IndividualTransport
}

func (t *TransportMeans1) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}
