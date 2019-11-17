package model

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters3 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max16Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max100KBinary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of parameters including the enveloppes) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType7 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters3) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters3) SetVersion(value string) {
	a.Version = (*Max16Text)(&value)
}

func (a *ApplicationParameters3) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max100KBinary)(&value))
}

func (a *ApplicationParameters3) AddEncryptedParameters() *ContentInformationType7 {
	a.EncryptedParameters = new(ContentInformationType7)
	return a.EncryptedParameters
}
