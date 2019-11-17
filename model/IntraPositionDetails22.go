package model

// Details of the intra-position movement.
type IntraPositionDetails22 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (e.g. restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`

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

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails22) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails22) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails22) AddCollateralMonitorAmount() *AmountAndDirection9 {
	i.CollateralMonitorAmount = new(AmountAndDirection9)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails22) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails22) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails22) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails22) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails22) AddCorporateActionEventType() *CorporateActionEventType3Choice {
	i.CorporateActionEventType = new(CorporateActionEventType3Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails22) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceFrom
}

func (i *IntraPositionDetails22) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceTo
}

func (i *IntraPositionDetails22) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
