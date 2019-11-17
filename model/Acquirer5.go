package model

// Acquirer involved in the card payment.
type Acquirer5 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification53 `xml:"Id"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max256Text `xml:"ParamsVrsn,omitempty"`
}

func (a *Acquirer5) AddIdentification() *GenericIdentification53 {
	a.Identification = new(GenericIdentification53)
	return a.Identification
}

func (a *Acquirer5) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max256Text)(&value)
}
