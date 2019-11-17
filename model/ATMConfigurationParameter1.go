package model

// Configuration parameter version of the ATM.
type ATMConfigurationParameter1 struct {

	// Type of the ATM configuration.
	Type *DataSetCategory7Code `xml:"Tp"`

	// Active version of the configuration.
	Version *Max35Text `xml:"Vrsn"`
}

func (a *ATMConfigurationParameter1) SetType(value string) {
	a.Type = (*DataSetCategory7Code)(&value)
}

func (a *ATMConfigurationParameter1) SetVersion(value string) {
	a.Version = (*Max35Text)(&value)
}
