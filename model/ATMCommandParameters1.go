package model

// Parameters to be used to update the configuration or the status security device.
type ATMCommandParameters1 struct {

	// Serial number of the device.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Update of the security configuration to apply on the security module of the ATM.
	RequiredConfiguration *ATMSecurityConfiguration1 `xml:"ReqrdCfgtn,omitempty"`

	// New status to apply on the security module of the ATM.
	RequiredStatus *ATMStatus2Code `xml:"ReqrdSts,omitempty"`
}

func (a *ATMCommandParameters1) SetSerialNumber(value string) {
	a.SerialNumber = (*Max35Text)(&value)
}

func (a *ATMCommandParameters1) AddRequiredConfiguration() *ATMSecurityConfiguration1 {
	a.RequiredConfiguration = new(ATMSecurityConfiguration1)
	return a.RequiredConfiguration
}

func (a *ATMCommandParameters1) SetRequiredStatus(value string) {
	a.RequiredStatus = (*ATMStatus2Code)(&value)
}
