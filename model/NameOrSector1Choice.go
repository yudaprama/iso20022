package model

// Provides the identification of the reported party through the name or the sector.
type NameOrSector1Choice struct {

	// Internal name of the counterparty of the reporting agent used by the reporting agent.
	Name *Max70Text `xml:"Nm"`

	// Represents the counterparty institutional section (such as non-financial corporation, central bank, ...).
	Sector *SNA2008SectorIdentifier `xml:"Sctr"`
}

func (n *NameOrSector1Choice) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}

func (n *NameOrSector1Choice) SetSector(value string) {
	n.Sector = (*SNA2008SectorIdentifier)(&value)
}
