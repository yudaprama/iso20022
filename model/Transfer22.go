package model

// Parameters applied to the settlement of a security transfer.
type Transfer22 struct {

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer22) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer22) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer22) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer22) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer22) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer22) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer22) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer22) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer22) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}
