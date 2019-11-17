package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.033.001.01 Document"`
	Message *SecuritiesFinancingInstructionV01 `xml:"SctiesFincgInstr"`
}

func (d *Document03300101) AddMessage() *SecuritiesFinancingInstructionV01 {
	d.Message = new(SecuritiesFinancingInstructionV01)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesFinancingInstruction to a securities financing transaction account servicer to notify the securities financing transaction account servicer of the details of a repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing transaction to allow the account servicer to manage the settlement and follow-up of the opening and closing leg of the transaction.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of securities financing transactions to a central securities depository or another settlement market infrastructure.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingInstructionV01 struct {

	// Information that unambiguously identifies a SecuritiesFinancingTransaction (unique per piece of collateral) and a SecuritiesFinancingInstruction message as known by the account owner (or the instructing party acting on its behalf).
	Identification *model.TransactionAndDocumentIdentification1 `xml:"Id"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionTypeAndAdditionalParameters *model.TransactionTypeAndAdditionalParameters1 `xml:"TxTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *model.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*model.Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *model.SecuritiesTradeDetails3 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount1 `xml:"QtyAndAcctDtls"`

	// Details for the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails1 `xml:"SctiesFincgDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails3 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction1 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties3 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *model.AmountAndDirection2 `xml:"OpngSttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts3 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties2 `xml:"OthrBizPties,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesFinancingInstructionV01) AddIdentification() *model.TransactionAndDocumentIdentification1 {
	s.Identification = new(model.TransactionAndDocumentIdentification1)
	return s.Identification
}

func (s *SecuritiesFinancingInstructionV01) AddTransactionTypeAndAdditionalParameters() *model.TransactionTypeAndAdditionalParameters1 {
	s.TransactionTypeAndAdditionalParameters = new(model.TransactionTypeAndAdditionalParameters1)
	return s.TransactionTypeAndAdditionalParameters
}

func (s *SecuritiesFinancingInstructionV01) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesFinancingInstructionV01) AddLinkages() *model.Linkages1 {
	newValue := new(model.Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesFinancingInstructionV01) AddTradeDetails() *model.SecuritiesTradeDetails3 {
	s.TradeDetails = new(model.SecuritiesTradeDetails3)
	return s.TradeDetails
}

func (s *SecuritiesFinancingInstructionV01) AddFinancialInstrumentIdentification() *model.SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingInstructionV01) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingInstructionV01) AddQuantityAndAccountDetails() *model.QuantityAndAccount1 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount1)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingInstructionV01) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails1 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails1)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingInstructionV01) AddSettlementParameters() *model.SettlementDetails3 {
	s.SettlementParameters = new(model.SettlementDetails3)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingInstructionV01) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction1 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction1)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingInstructionV01) AddDeliveringSettlementParties() *model.SettlementParties5 {
	s.DeliveringSettlementParties = new(model.SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingInstructionV01) AddReceivingSettlementParties() *model.SettlementParties5 {
	s.ReceivingSettlementParties = new(model.SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingInstructionV01) AddCashParties() *model.CashParties3 {
	s.CashParties = new(model.CashParties3)
	return s.CashParties
}

func (s *SecuritiesFinancingInstructionV01) AddOpeningSettlementAmount() *model.AmountAndDirection2 {
	s.OpeningSettlementAmount = new(model.AmountAndDirection2)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingInstructionV01) AddOtherAmounts() *model.OtherAmounts3 {
	s.OtherAmounts = new(model.OtherAmounts3)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingInstructionV01) AddOtherBusinessParties() *model.OtherParties2 {
	s.OtherBusinessParties = new(model.OtherParties2)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingInstructionV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesFinancingInstructionV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesFinancingInstructionV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
