package model

// Acquirer involved in the card payment.
type Acquirer4 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification53 `xml:"Id,omitempty"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max256Text `xml:"ParamsVrsn"`
}

func (a *Acquirer4) AddIdentification() *GenericIdentification53 {
	a.Identification = new(GenericIdentification53)
	return a.Identification
}

func (a *Acquirer4) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max256Text)(&value)
}
