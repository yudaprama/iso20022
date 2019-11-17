package model

// Acquirer involved in the card payment.
type Acquirer2 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max35Text `xml:"ParamsVrsn"`
}

func (a *Acquirer2) AddIdentification() *GenericIdentification32 {
	a.Identification = new(GenericIdentification32)
	return a.Identification
}

func (a *Acquirer2) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max35Text)(&value)
}
