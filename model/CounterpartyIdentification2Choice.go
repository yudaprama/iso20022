package model

// Provides the identification of the reporting agent counterparty.
type CounterpartyIdentification2Choice struct {

	// Legal entity identifier code used to recognise the counterparty of the reporting agent for the reported transaction.
	LEI *LEIIdentifier `xml:"LEI"`

	// Other identification of the counterparty through the sector and the location.
	Other *ReportedPartyIdentification1 `xml:"Othr"`
}

func (c *CounterpartyIdentification2Choice) SetLEI(value string) {
	c.LEI = (*LEIIdentifier)(&value)
}

func (c *CounterpartyIdentification2Choice) AddOther() *ReportedPartyIdentification1 {
	c.Other = new(ReportedPartyIdentification1)
	return c.Other
}
