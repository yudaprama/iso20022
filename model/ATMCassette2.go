package model

// Information on the cassette of an ATM.
type ATMCassette2 struct {

	// Physical identification of the cassette for the ATM.
	PhysicalIdentification *Max35Text `xml:"PhysId,omitempty"`

	// Logical identification of the cassette for the ATM.
	LogicalIdentification *Max35Text `xml:"LogclId"`

	// Serial number or unique identification of the cassette hardware.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Type of cassette.
	Type *ATMCassetteType1Code `xml:"Tp"`

	// Type of items the cash-in takes
	SubType []*ATMNoteType1Code `xml:"SubTp,omitempty"`

	// Type of media inside the cassette.
	MediaType *ATMMediaType1Code `xml:"MdiaTp,omitempty"`

	// Counter per unit value or globally.
	MediaCounters []*ATMCassetteCounters3 `xml:"MdiaCntrs,omitempty"`
}

func (a *ATMCassette2) SetPhysicalIdentification(value string) {
	a.PhysicalIdentification = (*Max35Text)(&value)
}

func (a *ATMCassette2) SetLogicalIdentification(value string) {
	a.LogicalIdentification = (*Max35Text)(&value)
}

func (a *ATMCassette2) SetSerialNumber(value string) {
	a.SerialNumber = (*Max35Text)(&value)
}

func (a *ATMCassette2) SetType(value string) {
	a.Type = (*ATMCassetteType1Code)(&value)
}

func (a *ATMCassette2) AddSubType(value string) {
	a.SubType = append(a.SubType, (*ATMNoteType1Code)(&value))
}

func (a *ATMCassette2) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType1Code)(&value)
}

func (a *ATMCassette2) AddMediaCounters() *ATMCassetteCounters3 {
	newValue := new(ATMCassetteCounters3)
	a.MediaCounters = append(a.MediaCounters, newValue)
	return newValue
}
