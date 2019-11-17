package model

// Content of the acceptor configuration.
type AcceptorConfigurationContent2 struct {

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters3 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*Max10000Binary `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*Max10000Binary `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters2 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host.
	HostCommunicationParameters []*HostCommunicationParameter2 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters1 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent2) AddAcquirerProtocolParameters() *AcquirerProtocolParameters3 {
	newValue := new(AcquirerProtocolParameters3)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent2) AddMerchantParameters(value string) {
	a.MerchantParameters = append(a.MerchantParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent2) AddTerminalParameters(value string) {
	a.TerminalParameters = append(a.TerminalParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent2) AddApplicationParameters() *ApplicationParameters2 {
	newValue := new(ApplicationParameters2)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent2) AddHostCommunicationParameters() *HostCommunicationParameter2 {
	newValue := new(HostCommunicationParameter2)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent2) AddSecurityParameters() *SecurityParameters1 {
	newValue := new(SecurityParameters1)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}
