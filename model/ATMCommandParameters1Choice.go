package model

// Specific parameters attached to an ATM command.
type ATMCommandParameters1Choice struct {

	// Required status of the ATM, parameters of the status update command.
	ATMRequiredGlobalStatus *ATMStatus1Code `xml:"ATMReqrdGblSts"`

	// Message to send for the send message command.
	ExpectedMessageFunction *MessageFunction8Code `xml:"XpctdMsgFctn"`

	// Parameters to be used by the configuration update command.
	RequiredConfigurationParameter *ATMConfigurationParameter1 `xml:"ReqrdCfgtnParam"`
}

func (a *ATMCommandParameters1Choice) SetATMRequiredGlobalStatus(value string) {
	a.ATMRequiredGlobalStatus = (*ATMStatus1Code)(&value)
}

func (a *ATMCommandParameters1Choice) SetExpectedMessageFunction(value string) {
	a.ExpectedMessageFunction = (*MessageFunction8Code)(&value)
}

func (a *ATMCommandParameters1Choice) AddRequiredConfigurationParameter() *ATMConfigurationParameter1 {
	a.RequiredConfigurationParameter = new(ATMConfigurationParameter1)
	return a.RequiredConfigurationParameter
}
