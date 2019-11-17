package model

// ATM terminal equipment.
type ATMEquipment1 struct {

	// ATM Manufacturer.
	Manufacturer *Max35Text `xml:"Manfctr,omitempty"`

	// Model of ATM.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Serial number of the ATM.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Provider of the ATM application software.
	ApplicationProvider *Max35Text `xml:"ApplPrvdr,omitempty"`

	// Name of the software product.
	ApplicationName *Max35Text `xml:"ApplNm,omitempty"`

	// Current version of the software that might include the release number.
	ApplicationVersion *Max35Text `xml:"ApplVrsn,omitempty"`

	// Unique assessment number for the component.
	ApprovalNumber *Max35Text `xml:"ApprvlNb,omitempty"`

	// Configuration parameter version.
	ConfigurationParameter []*ATMConfigurationParameter1 `xml:"CfgtnParam,omitempty"`
}

func (a *ATMEquipment1) SetManufacturer(value string) {
	a.Manufacturer = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetModel(value string) {
	a.Model = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetSerialNumber(value string) {
	a.SerialNumber = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApplicationProvider(value string) {
	a.ApplicationProvider = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApplicationName(value string) {
	a.ApplicationName = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApplicationVersion(value string) {
	a.ApplicationVersion = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApprovalNumber(value string) {
	a.ApprovalNumber = (*Max35Text)(&value)
}

func (a *ATMEquipment1) AddConfigurationParameter() *ATMConfigurationParameter1 {
	newValue := new(ATMConfigurationParameter1)
	a.ConfigurationParameter = append(a.ConfigurationParameter, newValue)
	return newValue
}
