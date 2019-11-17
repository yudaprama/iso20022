package model

// Choice between a corporate action identification or a corporate action official identification.
type CorporateActionEventReference3Choice struct {

	// Official and unique reference assigned by the official central body/ entity within each market at the beginning of a corporate action event.
	LinkedOfficialCorporateActionEventIdentification *Max35Text `xml:"LkdOffclCorpActnEvtId"`

	// Reference assigned by the account servicer to unambiguously identify a related corporate action event.
	LinkedCorporateActionIdentification *Max35Text `xml:"LkdCorpActnId"`
}

func (c *CorporateActionEventReference3Choice) SetLinkedOfficialCorporateActionEventIdentification(value string) {
	c.LinkedOfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionEventReference3Choice) SetLinkedCorporateActionIdentification(value string) {
	c.LinkedCorporateActionIdentification = (*Max35Text)(&value)
}
