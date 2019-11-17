package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02300101 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.023.001.01 Document"`
	Message *SecuritiesSettlementTransactionInstructionV01 `xml:"SctiesSttlmTxInstr"`
}

func (d *Document02300101) AddMessage() *SecuritiesSettlementTransactionInstructionV01 {
	d.Message = new(SecuritiesSettlementTransactionInstructionV01)
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
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of transactions to a central securities depository or another settlement market infrastructure
// Usage
// The instruction may be linked to other settlement instructions, for example, for a turnaround or back-to-back, or other transactions, for example, foreign exchange deal, using the linkage functionality.
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionInstructionV01 struct {

	// Information that unambiguously identifies a SecuritiesSettlementTransaction and a SecuritiesSettlementTransactionInstruction message as known by the account owner (or the instructing party acting on its behalf).
	Identification *model.TransactionAndDocumentIdentification1 `xml:"Id"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *model.SettlementTypeAndAdditionalParameters1 `xml:"SttlmTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *model.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*model.Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails1 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount1 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails1 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction1 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties3 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *model.AmountAndDirection2 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts3 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties2 `xml:"OthrBizPties,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddIdentification() *model.TransactionAndDocumentIdentification1 {
	s.Identification = new(model.TransactionAndDocumentIdentification1)
	return s.Identification
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddSettlementTypeAndAdditionalParameters() *model.SettlementTypeAndAdditionalParameters1 {
	s.SettlementTypeAndAdditionalParameters = new(model.SettlementTypeAndAdditionalParameters1)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddLinkages() *model.Linkages1 {
	newValue := new(model.Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddTradeDetails() *model.SecuritiesTradeDetails1 {
	s.TradeDetails = new(model.SecuritiesTradeDetails1)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddFinancialInstrumentIdentification() *model.SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddQuantityAndAccountDetails() *model.QuantityAndAccount1 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount1)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddSettlementParameters() *model.SettlementDetails1 {
	s.SettlementParameters = new(model.SettlementDetails1)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction1 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction1)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddDeliveringSettlementParties() *model.SettlementParties5 {
	s.DeliveringSettlementParties = new(model.SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddReceivingSettlementParties() *model.SettlementParties5 {
	s.ReceivingSettlementParties = new(model.SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddCashParties() *model.CashParties3 {
	s.CashParties = new(model.CashParties3)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddSettlementAmount() *model.AmountAndDirection2 {
	s.SettlementAmount = new(model.AmountAndDirection2)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddOtherAmounts() *model.OtherAmounts3 {
	s.OtherAmounts = new(model.OtherAmounts3)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddOtherBusinessParties() *model.OtherParties2 {
	s.OtherBusinessParties = new(model.OtherParties2)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionInstructionV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
