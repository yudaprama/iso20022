package model

// Content of the acceptor configuration.
type AcceptorConfigurationContent3 struct {

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters6 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*Max10000Binary `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*PaymentTerminalParameters1 `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters3 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host.
	HostCommunicationParameters []*HostCommunicationParameter2 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters2 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent3) AddAcquirerProtocolParameters() *AcquirerProtocolParameters6 {
	newValue := new(AcquirerProtocolParameters6)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddMerchantParameters(value string) {
	a.MerchantParameters = append(a.MerchantParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent3) AddTerminalParameters() *PaymentTerminalParameters1 {
	newValue := new(PaymentTerminalParameters1)
	a.TerminalParameters = append(a.TerminalParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddApplicationParameters() *ApplicationParameters3 {
	newValue := new(ApplicationParameters3)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddHostCommunicationParameters() *HostCommunicationParameter2 {
	newValue := new(HostCommunicationParameter2)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddSecurityParameters() *SecurityParameters2 {
	newValue := new(SecurityParameters2)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}
