package model

// Content of the acceptor configuration.
type AcceptorConfigurationContent1 struct {

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters1 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*Max10000Binary `xml:"MrchntParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters1 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host.
	HostCommunicationParameters []*HostCommunicationParameter1 `xml:"HstComParams,omitempty"`
}

func (a *AcceptorConfigurationContent1) AddAcquirerProtocolParameters() *AcquirerProtocolParameters1 {
	newValue := new(AcquirerProtocolParameters1)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent1) AddMerchantParameters(value string) {
	a.MerchantParameters = append(a.MerchantParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent1) AddApplicationParameters() *ApplicationParameters1 {
	newValue := new(ApplicationParameters1)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent1) AddHostCommunicationParameters() *HostCommunicationParameter1 {
	newValue := new(HostCommunicationParameter1)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}
