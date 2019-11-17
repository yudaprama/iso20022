package model

// Characteristics of a hardware memory module.
type MemoryCharacteristics1 struct {

	// Identification or name of the memory.
	Identification *Max35Text `xml:"Id"`

	// Total size of the memory unit.
	TotalSize *DecimalNumber `xml:"TtlSz"`

	// Total size of the available memory.
	FreeSize *DecimalNumber `xml:"FreeSz"`

	// Memory unit of the sizes.
	Unit *MemoryUnit1Code `xml:"Unit"`
}

func (m *MemoryCharacteristics1) SetIdentification(value string) {
	m.Identification = (*Max35Text)(&value)
}

func (m *MemoryCharacteristics1) SetTotalSize(value string) {
	m.TotalSize = (*DecimalNumber)(&value)
}

func (m *MemoryCharacteristics1) SetFreeSize(value string) {
	m.FreeSize = (*DecimalNumber)(&value)
}

func (m *MemoryCharacteristics1) SetUnit(value string) {
	m.Unit = (*MemoryUnit1Code)(&value)
}
