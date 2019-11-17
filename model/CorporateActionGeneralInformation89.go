package model

// General information about the corporate action event.
type CorporateActionGeneralInformation89 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType31Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat16Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity1Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation89) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation89) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation89) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation89) AddEventType() *CorporateActionEventType31Choice {
	c.EventType = new(CorporateActionEventType31Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation89) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation89) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat16Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat16Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation89) AddFractionalQuantity() *FinancialInstrumentQuantity1Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.FractionalQuantity
}
