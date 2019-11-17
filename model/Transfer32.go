package model

// Parameters applied to the settlement of a security transfer.
type Transfer32 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer32) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer32) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer32) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer32) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer32) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer32) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer32) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer32) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer32) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer32) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer32) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer32) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer32) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer32) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}
