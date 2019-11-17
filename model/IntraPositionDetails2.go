package model

// Details of the intra-position movement.
type IntraPositionDetails2 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType3Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Balance to which the securities were moved.
	BalanceTo *SecuritiesBalanceType3Choice `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails2) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails2) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails2) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails2) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails2) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails2) AddCorporateActionEventType() *CorporateActionEventType3Choice {
	i.CorporateActionEventType = new(CorporateActionEventType3Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails2) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails2) AddBalanceTo() *SecuritiesBalanceType3Choice {
	i.BalanceTo = new(SecuritiesBalanceType3Choice)
	return i.BalanceTo
}

func (i *IntraPositionDetails2) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
