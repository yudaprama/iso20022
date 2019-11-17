package model

// Specific parameters attached to an ATM command.
type ATMCommandParameters2Choice struct {

	// Required status of the ATM, parameters of the status update command.
	ATMRequiredGlobalStatus *ATMStatus1Code `xml:"ATMReqrdGblSts"`

	// Message to send for the send message command.
	ExpectedMessageFunction *MessageFunction8Code `xml:"XpctdMsgFctn"`

	// Parameters to be used by the configuration update command.
	RequiredConfigurationParameter *ATMConfigurationParameter1 `xml:"ReqrdCfgtnParam"`

	// Parameters to be used by the security scheme update command.
	RequiredSecurityScheme *ATMSecurityScheme2Code `xml:"ReqrdSctySchme"`

	// Parameters to be used to update the configuration or the status security device.
	SecurityDevice *ATMCommandParameters1 `xml:"SctyDvc"`

	// Parameters to be used by the various cryptographic key commands.
	Key *ATMConfigurationParameter2 `xml:"Key"`
}

func (a *ATMCommandParameters2Choice) SetATMRequiredGlobalStatus(value string) {
	a.ATMRequiredGlobalStatus = (*ATMStatus1Code)(&value)
}

func (a *ATMCommandParameters2Choice) SetExpectedMessageFunction(value string) {
	a.ExpectedMessageFunction = (*MessageFunction8Code)(&value)
}

func (a *ATMCommandParameters2Choice) AddRequiredConfigurationParameter() *ATMConfigurationParameter1 {
	a.RequiredConfigurationParameter = new(ATMConfigurationParameter1)
	return a.RequiredConfigurationParameter
}

func (a *ATMCommandParameters2Choice) SetRequiredSecurityScheme(value string) {
	a.RequiredSecurityScheme = (*ATMSecurityScheme2Code)(&value)
}

func (a *ATMCommandParameters2Choice) AddSecurityDevice() *ATMCommandParameters1 {
	a.SecurityDevice = new(ATMCommandParameters1)
	return a.SecurityDevice
}

func (a *ATMCommandParameters2Choice) AddKey() *ATMConfigurationParameter2 {
	a.Key = new(ATMConfigurationParameter2)
	return a.Key
}
