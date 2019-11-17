package model

// Context of the ATM for the key download.
type ATMSecurityContext2 struct {

	// Key exchange security scheme in activation on the ATM for the host manager.
	CurrentSecurityScheme *ATMSecurityScheme1Code `xml:"CurSctySchme"`

	// Hardware security module information, so called EPP for Encrypted PIN Pad.
	DeviceProperty *ATMEquipment3 `xml:"DvcPrprty,omitempty"`

	// Configuration parameters in use by the security device.
	CurrentConfiguration *ATMSecurityConfiguration1 `xml:"CurCfgtn,omitempty"`
}

func (a *ATMSecurityContext2) SetCurrentSecurityScheme(value string) {
	a.CurrentSecurityScheme = (*ATMSecurityScheme1Code)(&value)
}

func (a *ATMSecurityContext2) AddDeviceProperty() *ATMEquipment3 {
	a.DeviceProperty = new(ATMEquipment3)
	return a.DeviceProperty
}

func (a *ATMSecurityContext2) AddCurrentConfiguration() *ATMSecurityConfiguration1 {
	a.CurrentConfiguration = new(ATMSecurityConfiguration1)
	return a.CurrentConfiguration
}
