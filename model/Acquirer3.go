package model

// Acquirer involved in the card payment.
type Acquirer3 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification32 `xml:"Id"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max35Text `xml:"ParamsVrsn,omitempty"`
}

func (a *Acquirer3) AddIdentification() *GenericIdentification32 {
	a.Identification = new(GenericIdentification32)
	return a.Identification
}

func (a *Acquirer3) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max35Text)(&value)
}
