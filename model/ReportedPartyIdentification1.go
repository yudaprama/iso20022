package model

// Provides the identification of the reported party through the location and the name or the sector.
type ReportedPartyIdentification1 struct {

	// Name or sector of the counterparty of the reporting agent used by the reporting agent.
	NameOrSector *NameOrSector1Choice `xml:"NmOrSctr"`

	// Location of the country in which the counterparty is incorporated.
	Location *CountryCode `xml:"Lctn"`
}

func (r *ReportedPartyIdentification1) AddNameOrSector() *NameOrSector1Choice {
	r.NameOrSector = new(NameOrSector1Choice)
	return r.NameOrSector
}

func (r *ReportedPartyIdentification1) SetLocation(value string) {
	r.Location = (*CountryCode)(&value)
}
