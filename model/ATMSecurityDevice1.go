package model

// Hardware security module of the ATM.
type ATMSecurityDevice1 struct {

	// Hardware security module information, so called EPP for Encrypted PIN Pad.
	DeviceProperty *ATMEquipment2 `xml:"DvcPrprty,omitempty"`

	// Configuration parameters in use by the security device.
	CurrentConfiguration *ATMSecurityConfiguration1 `xml:"CurCfgtn"`

	// Configuration parameters supported by the security device.
	SupportedConfiguration *ATMSecurityConfiguration1 `xml:"SpprtdCfgtn,omitempty"`

	// Current status of the security device.
	CurrentStatus *ATMStatus2Code `xml:"CurSts"`

	// Incident occurring on the device.
	Incident *FailureReason5Code `xml:"Incdnt,omitempty"`
}

func (a *ATMSecurityDevice1) AddDeviceProperty() *ATMEquipment2 {
	a.DeviceProperty = new(ATMEquipment2)
	return a.DeviceProperty
}

func (a *ATMSecurityDevice1) AddCurrentConfiguration() *ATMSecurityConfiguration1 {
	a.CurrentConfiguration = new(ATMSecurityConfiguration1)
	return a.CurrentConfiguration
}

func (a *ATMSecurityDevice1) AddSupportedConfiguration() *ATMSecurityConfiguration1 {
	a.SupportedConfiguration = new(ATMSecurityConfiguration1)
	return a.SupportedConfiguration
}

func (a *ATMSecurityDevice1) SetCurrentStatus(value string) {
	a.CurrentStatus = (*ATMStatus2Code)(&value)
}

func (a *ATMSecurityDevice1) SetIncident(value string) {
	a.Incident = (*FailureReason5Code)(&value)
}
