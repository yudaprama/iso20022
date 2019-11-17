package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02300107 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.023.001.07 Document"`
	Message *SecuritiesSettlementTransactionInstructionV07 `xml:"SctiesSttlmTxInstr"`
}

func (d *Document02300107) AddMessage() *SecuritiesSettlementTransactionInstructionV07 {
	d.Message = new(SecuritiesSettlementTransactionInstructionV07)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementTransactionInstruction to an account servicer to instruct the receipt or delivery of financial instruments with or without payment, physically or by book-entry.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manages a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of transactions to a central securities depository or another settlement market infrastructure.
//
// Usage
// The instruction may be linked to other settlement instructions, for example, for a turnaround or back-to-back, or other transactions, for example, foreign exchange deal, using the linkage functionality.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionInstructionV07 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *model.SettlementTypeAndAdditionalParameters19 `xml:"SttlmTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *model.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*model.Linkages37 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails51 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount39 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails119 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *model.AmountAndDirection46 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionInstructionV07) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*model.Max35Text)(&value)
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddSettlementTypeAndAdditionalParameters() *model.SettlementTypeAndAdditionalParameters19 {
	s.SettlementTypeAndAdditionalParameters = new(model.SettlementTypeAndAdditionalParameters19)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddLinkages() *model.Linkages37 {
	newValue := new(model.Linkages37)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddTradeDetails() *model.SecuritiesTradeDetails51 {
	s.TradeDetails = new(model.SecuritiesTradeDetails51)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddFinancialInstrumentIdentification() *model.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddQuantityAndAccountDetails() *model.QuantityAndAccount39 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount39)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddSettlementParameters() *model.SettlementDetails119 {
	s.SettlementParameters = new(model.SettlementDetails119)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddDeliveringSettlementParties() *model.SettlementParties36 {
	s.DeliveringSettlementParties = new(model.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddReceivingSettlementParties() *model.SettlementParties36 {
	s.ReceivingSettlementParties = new(model.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddCashParties() *model.CashParties26 {
	s.CashParties = new(model.CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddSettlementAmount() *model.AmountAndDirection46 {
	s.SettlementAmount = new(model.AmountAndDirection46)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddOtherAmounts() *model.OtherAmounts28 {
	s.OtherAmounts = new(model.OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddOtherBusinessParties() *model.OtherParties27 {
	s.OtherBusinessParties = new(model.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionInstructionV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
