package model

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters4 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max100KBinary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of parameters including the envelope) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType10 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters4) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters4) SetVersion(value string) {
	a.Version = (*Max256Text)(&value)
}

func (a *ApplicationParameters4) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max100KBinary)(&value))
}

func (a *ApplicationParameters4) AddEncryptedParameters() *ContentInformationType10 {
	a.EncryptedParameters = new(ContentInformationType10)
	return a.EncryptedParameters
}
