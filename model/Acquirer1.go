package model

// Acquirer involved in the card payment.
type Acquirer1 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *ISODateTime `xml:"ParamsVrsn"`
}

func (a *Acquirer1) AddIdentification() *GenericIdentification32 {
	a.Identification = new(GenericIdentification32)
	return a.Identification
}

func (a *Acquirer1) SetParametersVersion(value string) {
	a.ParametersVersion = (*ISODateTime)(&value)
}
