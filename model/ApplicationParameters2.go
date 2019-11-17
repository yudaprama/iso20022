package model

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters2 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max16Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max10000Binary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of Parameters including the enveloppes) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType5 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters2) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters2) SetVersion(value string) {
	a.Version = (*Max16Text)(&value)
}

func (a *ApplicationParameters2) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max10000Binary)(&value))
}

func (a *ApplicationParameters2) AddEncryptedParameters() *ContentInformationType5 {
	a.EncryptedParameters = new(ContentInformationType5)
	return a.EncryptedParameters
}
