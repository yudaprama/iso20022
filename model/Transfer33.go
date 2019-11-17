package model

// Parameters applied to the settlement of a security transfer.
type Transfer33 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Specifies the reason for the assets transfer.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit6 `xml:"UnitsDtls,omitempty"`

	// Weighted average price of the units in the account before the transfer was executed.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Weighted average price of the units in the account after the transfer was executed.
	NewAveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"NewAvrgPric,omitempty"`

	// Trade date of the average weighted data of units in the account before the transfer was executed.
	AverageDate *ISODate `xml:"AvrgDt,omitempty"`

	// Trade date of the average weighted data of units in the account after the transfer was executed.
	NewAverageDate *ISODate `xml:"NewAvrgDt,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer33) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer33) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer33) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer33) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer33) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer33) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer33) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer33) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer33) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer33) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer33) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer33) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer33) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer33) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer33) AddUnitsDetails() *Unit6 {
	newValue := new(Unit6)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer33) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer33) SetNewAveragePrice(value, currency string) {
	t.NewAveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer33) SetAverageDate(value string) {
	t.AverageDate = (*ISODate)(&value)
}

func (t *Transfer33) SetNewAverageDate(value string) {
	t.NewAverageDate = (*ISODate)(&value)
}

func (t *Transfer33) SetTransferCurrency(value string) {
	t.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (t *Transfer33) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer33) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer33) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer33) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}

func (t *Transfer33) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}
