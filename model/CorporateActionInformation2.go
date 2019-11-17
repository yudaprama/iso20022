package model

// Provides information about the corporate action event.
type CorporateActionInformation2 struct {

	// Identification of the issuer's agent or the issuer.
	AgentIdentification *PartyIdentification2Choice `xml:"AgtId"`

	// Reference given to the event by the CA event issuer (agent).
	IssuerCorporateActionIdentification *Max35Text `xml:"IssrCorpActnId"`

	// Reference assigned by the (I)CSD to unambiguously identify a corporate avent.
	CorporateActionProcessingIdentification *Max35Text `xml:"CorpActnPrcgId,omitempty"`

	// Speficies the type of corporate event.
	EventType *CorporateActionEventType2FormatChoice `xml:"EvtTp"`

	// Type of processing involved by a Corporate Action.
	EventProcessingType *CorporateActionEventProcessingType1FormatChoice `xml:"EvtPrcgTp,omitempty"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1FormatChoice `xml:"MndtryVlntryEvtTp"`

	// Identification of the underlying financial instrument, ie, the financial instrument affected by the corporate action event.
	UnderlyingSecurity *FinancialInstrumentDescription3 `xml:"UndrlygScty"`

	// Identification of the secondary underlying financial instrument, ie, the non-principal financial instrument affected by the corporate action event.
	OtherUnderlyingSecurity []*FinancialInstrumentDescription3 `xml:"OthrUndrlygScty,omitempty"`
}

func (c *CorporateActionInformation2) AddAgentIdentification() *PartyIdentification2Choice {
	c.AgentIdentification = new(PartyIdentification2Choice)
	return c.AgentIdentification
}

func (c *CorporateActionInformation2) SetIssuerCorporateActionIdentification(value string) {
	c.IssuerCorporateActionIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionInformation2) SetCorporateActionProcessingIdentification(value string) {
	c.CorporateActionProcessingIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionInformation2) AddEventType() *CorporateActionEventType2FormatChoice {
	c.EventType = new(CorporateActionEventType2FormatChoice)
	return c.EventType
}

func (c *CorporateActionInformation2) AddEventProcessingType() *CorporateActionEventProcessingType1FormatChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingType1FormatChoice)
	return c.EventProcessingType
}

func (c *CorporateActionInformation2) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1FormatChoice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1FormatChoice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionInformation2) AddUnderlyingSecurity() *FinancialInstrumentDescription3 {
	c.UnderlyingSecurity = new(FinancialInstrumentDescription3)
	return c.UnderlyingSecurity
}

func (c *CorporateActionInformation2) AddOtherUnderlyingSecurity() *FinancialInstrumentDescription3 {
	newValue := new(FinancialInstrumentDescription3)
	c.OtherUnderlyingSecurity = append(c.OtherUnderlyingSecurity, newValue)
	return newValue
}
