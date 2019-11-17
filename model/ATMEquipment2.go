package model

// Hardware security module information, so called EPP for Encrypted PIN Pad.
type ATMEquipment2 struct {

	// ATM Manufacturer.
	Manufacturer *Max35Text `xml:"Manfctr,omitempty"`

	// Model of ATM.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Version of the device model.
	Version *Max35Text `xml:"Vrsn,omitempty"`

	// Serial number of the ATM.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Provider of the firmware.
	FirmwareProvider *Max35Text `xml:"FrmwrPrvdr,omitempty"`

	// Identification of the firmware.
	FirmwareIdentification *Max35Text `xml:"FrmwrId,omitempty"`

	// Version of the firmware.
	FirmwareVersion *Max35Text `xml:"FrmwrVrsn,omitempty"`
}

func (a *ATMEquipment2) SetManufacturer(value string) {
	a.Manufacturer = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetModel(value string) {
	a.Model = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetVersion(value string) {
	a.Version = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetSerialNumber(value string) {
	a.SerialNumber = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetFirmwareProvider(value string) {
	a.FirmwareProvider = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetFirmwareIdentification(value string) {
	a.FirmwareIdentification = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetFirmwareVersion(value string) {
	a.FirmwareVersion = (*Max35Text)(&value)
}
