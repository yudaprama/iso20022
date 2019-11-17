package model

// Choice between a corporate action identification or a corporate action official identification.
type CorporateActionEventReference4Choice struct {

	// Official and unique reference assigned by the official central body/ entity within each market at the beginning of a corporate action event.
	LinkedOfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"LkdOffclCorpActnEvtId"`

	// Reference assigned by the account servicer to unambiguously identify a related corporate action event.
	LinkedCorporateActionIdentification *RestrictedFINXMax16Text `xml:"LkdCorpActnId"`
}

func (c *CorporateActionEventReference4Choice) SetLinkedOfficialCorporateActionEventIdentification(value string) {
	c.LinkedOfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionEventReference4Choice) SetLinkedCorporateActionIdentification(value string) {
	c.LinkedCorporateActionIdentification = (*RestrictedFINXMax16Text)(&value)
}
