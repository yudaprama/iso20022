package model

// Information on the cassette of an ATM.
type ATMCassette1 struct {

	// Physical identification of the cassette for the ATM.
	PhysicalIdentification *Max35Text `xml:"PhysId,omitempty"`

	// Logical identification of the cassette for the ATM.
	LogicalIdentification *Max35Text `xml:"LogclId"`

	// Type of cassette.
	Type *ATMCassetteType1Code `xml:"Tp"`

	// Type of items the cash-in takes
	SubType []*ATMNoteType1Code `xml:"SubTp,omitempty"`

	// Type of media inside the cassette.
	MediaType *ATMMediaType1Code `xml:"MdiaTp,omitempty"`

	// Counter per unit value or globally.
	MediaCounters []*ATMCassetteCounters1 `xml:"MdiaCntrs,omitempty"`
}

func (a *ATMCassette1) SetPhysicalIdentification(value string) {
	a.PhysicalIdentification = (*Max35Text)(&value)
}

func (a *ATMCassette1) SetLogicalIdentification(value string) {
	a.LogicalIdentification = (*Max35Text)(&value)
}

func (a *ATMCassette1) SetType(value string) {
	a.Type = (*ATMCassetteType1Code)(&value)
}

func (a *ATMCassette1) AddSubType(value string) {
	a.SubType = append(a.SubType, (*ATMNoteType1Code)(&value))
}

func (a *ATMCassette1) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType1Code)(&value)
}

func (a *ATMCassette1) AddMediaCounters() *ATMCassetteCounters1 {
	newValue := new(ATMCassetteCounters1)
	a.MediaCounters = append(a.MediaCounters, newValue)
	return newValue
}
