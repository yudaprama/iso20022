package model

// Provides information on the settlement of a treasury trade.
type SettlementData2 struct {

	// Unique reference supplied by the trade processing system.
	CashFlowUniqueReference *Max35Text `xml:"CshFlowUnqRef,omitempty"`

	// Unique reference assigned by a settlement system.
	SettlementSystemUniqueReference *Max35Text `xml:"SttlmSysUnqRef,omitempty"`

	// Original amount which should be settled. This information should be provided when the trade is partially settled or when the settlement is rejected.
	SettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SttlmAmt"`

	// Funds which the trading side is expected to receive.
	SettledAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SttldAmt,omitempty"`

	// Amount that cannot be settled by a settlement system.
	RejectedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RjctdAmt,omitempty"`

	// Specifies the party that pays the settlement amount.
	PayingParty *PartyIdentification8Choice `xml:"PngPty"`

	// Specifies the party that receives the settlement amount.
	ReceivingParty *PartyIdentification8Choice `xml:"RcvgPty"`

	// Date on which the settlement is due to settle.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Specifies the status of a settlement eg rejected, settled or awaiting authorisation.
	SettlementStatus *SettlementStatus1Code `xml:"SttlmSts"`

	// Description of the status of the settlement of a trade when no coded form is available.
	ExtendedSettlementStatus *Extended350Code `xml:"XtndedSttlmSts"`

	// Additional information about the cause of the rejection of a settlement.
	SettlementStatusSubType *Max70Text `xml:"SttlmStsSubTp,omitempty"`

	// Cash settlement is suspended.
	Suspended *YesNoIndicator `xml:"Sspd"`

	// Cash settlement is pending.
	Pending *YesNoIndicator `xml:"Pdg"`
}

func (s *SettlementData2) SetCashFlowUniqueReference(value string) {
	s.CashFlowUniqueReference = (*Max35Text)(&value)
}

func (s *SettlementData2) SetSettlementSystemUniqueReference(value string) {
	s.SettlementSystemUniqueReference = (*Max35Text)(&value)
}

func (s *SettlementData2) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SettlementData2) SetSettledAmount(value, currency string) {
	s.SettledAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SettlementData2) SetRejectedAmount(value, currency string) {
	s.RejectedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SettlementData2) AddPayingParty() *PartyIdentification8Choice {
	s.PayingParty = new(PartyIdentification8Choice)
	return s.PayingParty
}

func (s *SettlementData2) AddReceivingParty() *PartyIdentification8Choice {
	s.ReceivingParty = new(PartyIdentification8Choice)
	return s.ReceivingParty
}

func (s *SettlementData2) SetSettlementDate(value string) {
	s.SettlementDate = (*ISODate)(&value)
}

func (s *SettlementData2) SetSettlementStatus(value string) {
	s.SettlementStatus = (*SettlementStatus1Code)(&value)
}

func (s *SettlementData2) SetExtendedSettlementStatus(value string) {
	s.ExtendedSettlementStatus = (*Extended350Code)(&value)
}

func (s *SettlementData2) SetSettlementStatusSubType(value string) {
	s.SettlementStatusSubType = (*Max70Text)(&value)
}

func (s *SettlementData2) SetSuspended(value string) {
	s.Suspended = (*YesNoIndicator)(&value)
}

func (s *SettlementData2) SetPending(value string) {
	s.Pending = (*YesNoIndicator)(&value)
}
