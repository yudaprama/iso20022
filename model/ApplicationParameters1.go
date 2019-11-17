package model

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters1 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max16Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max10000Binary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of Parameters including the enveloppes) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType2 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters1) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters1) SetVersion(value string) {
	a.Version = (*Max16Text)(&value)
}

func (a *ApplicationParameters1) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max10000Binary)(&value))
}

func (a *ApplicationParameters1) AddEncryptedParameters() *ContentInformationType2 {
	a.EncryptedParameters = new(ContentInformationType2)
	return a.EncryptedParameters
}
