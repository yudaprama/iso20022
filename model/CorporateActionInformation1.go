package model

// General information about the corporate action event.
type CorporateActionInformation1 struct {

	// Identification of the issuer agent.
	AgentIdentification *PartyIdentification2Choice `xml:"AgtId"`

	// Reference given to the event by the CA event issuer (agent).
	IssuerCorporateActionIdentification *Max35Text `xml:"IssrCorpActnId,omitempty"`

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionProcessingIdentification *Max35Text `xml:"CorpActnPrcgId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType2FormatChoice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1FormatChoice `xml:"MndtryVlntryEvtTp"`

	// Type of processing involved by a Corporate Action.
	EventProcessingType *CorporateActionEventProcessingType1FormatChoice `xml:"EvtPrcgTp,omitempty"`

	// Identification of the underlying financial instrument, ie, the financial instrument affected by the corporate action event.
	UnderlyingSecurity *FinancialInstrumentDescription3 `xml:"UndrlygScty"`
}

func (c *CorporateActionInformation1) AddAgentIdentification() *PartyIdentification2Choice {
	c.AgentIdentification = new(PartyIdentification2Choice)
	return c.AgentIdentification
}

func (c *CorporateActionInformation1) SetIssuerCorporateActionIdentification(value string) {
	c.IssuerCorporateActionIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionInformation1) SetCorporateActionProcessingIdentification(value string) {
	c.CorporateActionProcessingIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionInformation1) AddEventType() *CorporateActionEventType2FormatChoice {
	c.EventType = new(CorporateActionEventType2FormatChoice)
	return c.EventType
}

func (c *CorporateActionInformation1) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1FormatChoice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1FormatChoice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionInformation1) AddEventProcessingType() *CorporateActionEventProcessingType1FormatChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingType1FormatChoice)
	return c.EventProcessingType
}

func (c *CorporateActionInformation1) AddUnderlyingSecurity() *FinancialInstrumentDescription3 {
	c.UnderlyingSecurity = new(FinancialInstrumentDescription3)
	return c.UnderlyingSecurity
}
