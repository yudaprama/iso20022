package model

// Provides information about the standing instruction.
type CorporateActionStandingInstructionGeneralInformation1 struct {

	// Type of standing instruction.
	StandingInstructionType *StandingInstructionType1Code `xml:"StgInstrTp"`

	// Type of coporpate action event.
	EventType []*CorporateActionEventType2FormatChoice `xml:"EvtTp,omitempty"`

	// Identification of the instructing party, ie, the CSD client.
	InstructingPartyIdentification *PartyIdentification2Choice `xml:"InstgPtyId"`

	// Reference of the standing instruction assigned by the client.
	ClientStandingInstructionIdentification *Max35Text `xml:"ClntStgInstrId"`

	// Provides information about the account to which the standing instruction can apply.
	AccountDetails []*IncludedAccount1 `xml:"AcctDtls,omitempty"`

	// Identification of the underlying financial instrument, ie, the financial instrument affected by the corporate action event.
	UnderlyingSecurity *FinancialInstrumentDescription3 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionStandingInstructionGeneralInformation1) SetStandingInstructionType(value string) {
	c.StandingInstructionType = (*StandingInstructionType1Code)(&value)
}

func (c *CorporateActionStandingInstructionGeneralInformation1) AddEventType() *CorporateActionEventType2FormatChoice {
	newValue := new(CorporateActionEventType2FormatChoice)
	c.EventType = append(c.EventType, newValue)
	return newValue
}

func (c *CorporateActionStandingInstructionGeneralInformation1) AddInstructingPartyIdentification() *PartyIdentification2Choice {
	c.InstructingPartyIdentification = new(PartyIdentification2Choice)
	return c.InstructingPartyIdentification
}

func (c *CorporateActionStandingInstructionGeneralInformation1) SetClientStandingInstructionIdentification(value string) {
	c.ClientStandingInstructionIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionStandingInstructionGeneralInformation1) AddAccountDetails() *IncludedAccount1 {
	newValue := new(IncludedAccount1)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}

func (c *CorporateActionStandingInstructionGeneralInformation1) AddUnderlyingSecurity() *FinancialInstrumentDescription3 {
	c.UnderlyingSecurity = new(FinancialInstrumentDescription3)
	return c.UnderlyingSecurity
}
