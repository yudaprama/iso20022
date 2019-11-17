package model

// Content of the acceptor configuration.
type AcceptorConfigurationContent5 struct {

	// True if the whole configuration related to the terminal manager has to be replaced by the configuration included in the message content.
	ReplaceConfiguration *TrueFalseIndicator `xml:"RplcCfgtn,omitempty"`

	// Configuration parameters of the TMS protocol between a POI and a terminal manager.
	TMSProtocolParameters []*TMSProtocolParameters2 `xml:"TMSPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters9 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*MerchantConfigurationParameters2 `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*PaymentTerminalParameters3 `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters5 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host or a terminal manager host.
	HostCommunicationParameters []*HostCommunicationParameter4 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters6 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent5) SetReplaceConfiguration(value string) {
	a.ReplaceConfiguration = (*TrueFalseIndicator)(&value)
}

func (a *AcceptorConfigurationContent5) AddTMSProtocolParameters() *TMSProtocolParameters2 {
	newValue := new(TMSProtocolParameters2)
	a.TMSProtocolParameters = append(a.TMSProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddAcquirerProtocolParameters() *AcquirerProtocolParameters9 {
	newValue := new(AcquirerProtocolParameters9)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddMerchantParameters() *MerchantConfigurationParameters2 {
	newValue := new(MerchantConfigurationParameters2)
	a.MerchantParameters = append(a.MerchantParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddTerminalParameters() *PaymentTerminalParameters3 {
	newValue := new(PaymentTerminalParameters3)
	a.TerminalParameters = append(a.TerminalParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddApplicationParameters() *ApplicationParameters5 {
	newValue := new(ApplicationParameters5)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddHostCommunicationParameters() *HostCommunicationParameter4 {
	newValue := new(HostCommunicationParameter4)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddSecurityParameters() *SecurityParameters6 {
	newValue := new(SecurityParameters6)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}
