package model

// Describes the multimodal or the individual transport of goods.
type TransportMeans5 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport7 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans5) AddIndividualTransport() *SingleTransport7 {
	t.IndividualTransport = new(SingleTransport7)
	return t.IndividualTransport
}

func (t *TransportMeans5) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}
