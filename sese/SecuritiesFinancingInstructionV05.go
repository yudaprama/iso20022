package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300105 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.033.001.05 Document"`
	Message *SecuritiesFinancingInstructionV05 `xml:"SctiesFincgInstr"`
}

func (d *Document03300105) AddMessage() *SecuritiesFinancingInstructionV05 {
	d.Message = new(SecuritiesFinancingInstructionV05)
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
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingInstructionV05 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionTypeAndAdditionalParameters *model.TransactionTypeAndAdditionalParameters1 `xml:"TxTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *model.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*model.Linkages17 `xml:"Lnkgs,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *model.SecuritiesTradeDetails37 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount25 `xml:"QtyAndAcctDtls"`

	// Details for the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails11 `xml:"SctiesFincgDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails72 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction3 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *model.AmountAndDirection36 `xml:"OpngSttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingInstructionV05) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*model.Max35Text)(&value)
}

func (s *SecuritiesFinancingInstructionV05) AddTransactionTypeAndAdditionalParameters() *model.TransactionTypeAndAdditionalParameters1 {
	s.TransactionTypeAndAdditionalParameters = new(model.TransactionTypeAndAdditionalParameters1)
	return s.TransactionTypeAndAdditionalParameters
}

func (s *SecuritiesFinancingInstructionV05) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesFinancingInstructionV05) AddLinkages() *model.Linkages17 {
	newValue := new(model.Linkages17)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesFinancingInstructionV05) AddTradeDetails() *model.SecuritiesTradeDetails37 {
	s.TradeDetails = new(model.SecuritiesTradeDetails37)
	return s.TradeDetails
}

func (s *SecuritiesFinancingInstructionV05) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingInstructionV05) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingInstructionV05) AddQuantityAndAccountDetails() *model.QuantityAndAccount25 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount25)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingInstructionV05) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails11 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails11)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingInstructionV05) AddSettlementParameters() *model.SettlementDetails72 {
	s.SettlementParameters = new(model.SettlementDetails72)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingInstructionV05) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction3 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction3)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingInstructionV05) AddDeliveringSettlementParties() *model.SettlementParties10 {
	s.DeliveringSettlementParties = new(model.SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingInstructionV05) AddReceivingSettlementParties() *model.SettlementParties10 {
	s.ReceivingSettlementParties = new(model.SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingInstructionV05) AddCashParties() *model.CashParties17 {
	s.CashParties = new(model.CashParties17)
	return s.CashParties
}

func (s *SecuritiesFinancingInstructionV05) AddOpeningSettlementAmount() *model.AmountAndDirection36 {
	s.OpeningSettlementAmount = new(model.AmountAndDirection36)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingInstructionV05) AddOtherAmounts() *model.OtherAmounts14 {
	s.OtherAmounts = new(model.OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingInstructionV05) AddOtherBusinessParties() *model.OtherParties19 {
	s.OtherBusinessParties = new(model.OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingInstructionV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
