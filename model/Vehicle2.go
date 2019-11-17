package model

// Additional information related to a vehicle.
type Vehicle2 struct {

	// Type of information related to the vehicle.
	Type *Max35Text `xml:"Tp,omitempty"`

	// Entry mode of the information.
	EntryMode *CardDataReading5Code `xml:"NtryMd,omitempty"`

	// Information related to the vehicle.
	Data *Max35Text `xml:"Data"`
}

func (v *Vehicle2) SetType(value string) {
	v.Type = (*Max35Text)(&value)
}

func (v *Vehicle2) SetEntryMode(value string) {
	v.EntryMode = (*CardDataReading5Code)(&value)
}

func (v *Vehicle2) SetData(value string) {
	v.Data = (*Max35Text)(&value)
}
