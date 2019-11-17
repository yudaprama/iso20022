package model

// Describes the multimodal or the individual transport of goods.
type TransportMeans2 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport5 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans2) AddIndividualTransport() *SingleTransport5 {
	t.IndividualTransport = new(SingleTransport5)
	return t.IndividualTransport
}

func (t *TransportMeans2) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}
