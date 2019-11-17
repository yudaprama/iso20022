package model

// Content of the acceptor configuration.
type AcceptorConfigurationContent4 struct {

	// Configuration parameters of the TMS protocol between a POI and a terminal manager.
	TMSProtocolParameters []*TMSProtocolParameters1 `xml:"TMSPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters7 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*MerchantConfigurationParameters1 `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*PaymentTerminalParameters2 `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters4 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host or a terminal manager host.
	HostCommunicationParameters []*HostCommunicationParameter3 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters3 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent4) AddTMSProtocolParameters() *TMSProtocolParameters1 {
	newValue := new(TMSProtocolParameters1)
	a.TMSProtocolParameters = append(a.TMSProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddAcquirerProtocolParameters() *AcquirerProtocolParameters7 {
	newValue := new(AcquirerProtocolParameters7)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddMerchantParameters() *MerchantConfigurationParameters1 {
	newValue := new(MerchantConfigurationParameters1)
	a.MerchantParameters = append(a.MerchantParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddTerminalParameters() *PaymentTerminalParameters2 {
	newValue := new(PaymentTerminalParameters2)
	a.TerminalParameters = append(a.TerminalParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddApplicationParameters() *ApplicationParameters4 {
	newValue := new(ApplicationParameters4)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddHostCommunicationParameters() *HostCommunicationParameter3 {
	newValue := new(HostCommunicationParameter3)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddSecurityParameters() *SecurityParameters3 {
	newValue := new(SecurityParameters3)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}
