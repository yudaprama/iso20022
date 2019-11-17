package model

// Details of the intra-position movement.
type IntraPositionMovementDetails1 struct {

	// Identifications (account owner and/or account servicer) of the intra-position movement.
	Identification *References5Choice `xml:"Id,omitempty"`

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Balance to which the securities were moved.
	BalanceTo *SecuritiesBalanceType3Choice `xml:"BalTo"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType3Choice `xml:"CorpActnEvtTp,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (i *IntraPositionMovementDetails1) AddIdentification() *References5Choice {
	i.Identification = new(References5Choice)
	return i.Identification
}

func (i *IntraPositionMovementDetails1) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionMovementDetails1) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionMovementDetails1) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionMovementDetails1) AddBalanceTo() *SecuritiesBalanceType3Choice {
	i.BalanceTo = new(SecuritiesBalanceType3Choice)
	return i.BalanceTo
}

func (i *IntraPositionMovementDetails1) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionMovementDetails1) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionMovementDetails1) AddCorporateActionEventType() *CorporateActionEventType3Choice {
	i.CorporateActionEventType = new(CorporateActionEventType3Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionMovementDetails1) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (i *IntraPositionMovementDetails1) AddExtension() *Extension2 {
	newValue := new(Extension2)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
