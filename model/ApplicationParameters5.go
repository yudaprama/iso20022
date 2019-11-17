package model

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters5 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max100KBinary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of parameters including the envelope) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType10 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters5) SetActionType(value string) {
	a.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (a *ApplicationParameters5) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters5) SetVersion(value string) {
	a.Version = (*Max256Text)(&value)
}

func (a *ApplicationParameters5) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max100KBinary)(&value))
}

func (a *ApplicationParameters5) AddEncryptedParameters() *ContentInformationType10 {
	a.EncryptedParameters = new(ContentInformationType10)
	return a.EncryptedParameters
}
