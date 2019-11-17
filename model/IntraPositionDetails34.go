package model

// Details of the intra-position movement.
type IntraPositionDetails34 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection44 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType29Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails34) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails34) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails34) AddCollateralMonitorAmount() *AmountAndDirection44 {
	i.CollateralMonitorAmount = new(AmountAndDirection44)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails34) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails34) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails34) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails34) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails34) AddCorporateActionEventType() *CorporateActionEventType29Choice {
	i.CorporateActionEventType = new(CorporateActionEventType29Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails34) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceFrom
}

func (i *IntraPositionDetails34) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceTo
}

func (i *IntraPositionDetails34) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
