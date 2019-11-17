package model

// Provides the clearing details.
type Clearing3 struct {

	// Provides details about the clearing member identification and account.
	ClearingMember []*PartyIdentificationAndAccount78 `xml:"ClrMmb"`

	// Clearing organisation that will clear the trade.
	// Note: This field allows Clearing Member Firm to segregate flows coming from clearing counterparty's clearing system. Indeed, Clearing Member Firms receive messages from the same system (same sender) and this field allows them to know if the message is related to equities or derivatives.
	ClearingSegment *PartyIdentification35Choice `xml:"ClrSgmt,omitempty"`
}

func (c *Clearing3) AddClearingMember() *PartyIdentificationAndAccount78 {
	newValue := new(PartyIdentificationAndAccount78)
	c.ClearingMember = append(c.ClearingMember, newValue)
	return newValue
}

func (c *Clearing3) AddClearingSegment() *PartyIdentification35Choice {
	c.ClearingSegment = new(PartyIdentification35Choice)
	return c.ClearingSegment
}
