package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300206 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.033.002.06 Document"`
	Message *SecuritiesFinancingInstruction002V06 `xml:"SctiesFincgInstr"`
}

func (d *Document03300206) AddMessage() *SecuritiesFinancingInstruction002V06 {
	d.Message = new(SecuritiesFinancingInstruction002V06)
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
type SecuritiesFinancingInstruction002V06 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.RestrictedFINXMax16Text `xml:"TxId"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionTypeAndAdditionalParameters *model.TransactionTypeAndAdditionalParameters13 `xml:"TxTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *model.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*model.Linkages43 `xml:"Lnkgs,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *model.SecuritiesTradeDetails59 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount47 `xml:"QtyAndAcctDtls"`

	// Details for the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails30 `xml:"SctiesFincgDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails105 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *model.AmountAndDirection60 `xml:"OpngSttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingInstruction002V06) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*model.RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingInstruction002V06) AddTransactionTypeAndAdditionalParameters() *model.TransactionTypeAndAdditionalParameters13 {
	s.TransactionTypeAndAdditionalParameters = new(model.TransactionTypeAndAdditionalParameters13)
	return s.TransactionTypeAndAdditionalParameters
}

func (s *SecuritiesFinancingInstruction002V06) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesFinancingInstruction002V06) AddLinkages() *model.Linkages43 {
	newValue := new(model.Linkages43)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesFinancingInstruction002V06) AddTradeDetails() *model.SecuritiesTradeDetails59 {
	s.TradeDetails = new(model.SecuritiesTradeDetails59)
	return s.TradeDetails
}

func (s *SecuritiesFinancingInstruction002V06) AddFinancialInstrumentIdentification() *model.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingInstruction002V06) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingInstruction002V06) AddQuantityAndAccountDetails() *model.QuantityAndAccount47 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount47)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingInstruction002V06) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails30 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails30)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingInstruction002V06) AddSettlementParameters() *model.SettlementDetails105 {
	s.SettlementParameters = new(model.SettlementDetails105)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingInstruction002V06) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingInstruction002V06) AddDeliveringSettlementParties() *model.SettlementParties44 {
	s.DeliveringSettlementParties = new(model.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingInstruction002V06) AddReceivingSettlementParties() *model.SettlementParties44 {
	s.ReceivingSettlementParties = new(model.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingInstruction002V06) AddCashParties() *model.CashParties30 {
	s.CashParties = new(model.CashParties30)
	return s.CashParties
}

func (s *SecuritiesFinancingInstruction002V06) AddOpeningSettlementAmount() *model.AmountAndDirection60 {
	s.OpeningSettlementAmount = new(model.AmountAndDirection60)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingInstruction002V06) AddOtherAmounts() *model.OtherAmounts35 {
	s.OtherAmounts = new(model.OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingInstruction002V06) AddOtherBusinessParties() *model.OtherParties29 {
	s.OtherBusinessParties = new(model.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingInstruction002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
