package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02700101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.027.001.01 Document"`
	Message *SecuritiesTradeConfirmationV01 `xml:"SctiesTradConf"`
}

func (d *Document02700101) AddMessage() *SecuritiesTradeConfirmationV01 {
	d.Message = new(SecuritiesTradeConfirmationV01)
	return d.Message
}

// Scope
// Sent by an executing party to an instructing party directly or through Central Matching Utility (CMU) to provide trade confirmation on a per-account basis based on instructions provided by the instructing party in the SecuritiesAllocationInstruction message.
// It may also be used to provide trade confirmation on the trade level from an executing party or an instructing party to the custodian or an affirming party directly or through CMU.
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager or of other categories of investors.
// The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// The custodian or the affirming party is typically the custodian, trustee, financial institution, intermediary system/vendor communicating on behalf of them, or their agent.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: In local matching, the initiator of this message is always the executing party. In central matching the initiator may be either the executing party or instructing party.
// Respondent: instructing party, a custodian or an affirming party responds with SecuritiesTradeConfirmationResponse (accept or reject) message.
type SecuritiesTradeConfirmationV01 struct {

	// Information that unambiguously identifies an SecuritiesTradeConfirmation message as known by the account owner (or the instructing party acting on its behalf).
	Identification *model.TransactiontIdentification4 `xml:"Id"`

	// Count of the number of transactions linked.
	NumberCount *model.NumberCount1Choice `xml:"NbCnt,omitempty"`

	// Reference to the transaction identifier issued by a business party and/or market infrastructure. It may also be used to reference a previous transaction, for example, a block/allocation instruction, or tie a set of messages together.
	References []*model.Linkages15 `xml:"Refs,omitempty"`

	// Details of the trade.
	TradeDetails *model.Order14 `xml:"TradDtls"`

	// Unique and unambiguous identifier of a financial instrument, assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes31 `xml:"FinInstrmAttrbts,omitempty"`

	// Underlying financial instrument to which an trade confirmation is related.
	UnderlyingFinancialInstrument []*model.UnderlyingFinancialInstrument1 `xml:"UndrlygFinInstrm,omitempty"`

	// Additional restrictions on the financial instrument, related to the stipulation.
	Stipulations *model.FinancialInstrumentStipulations2 `xml:"Stiptns,omitempty"`

	// Parties involved in the confirmation of the details of a trade.
	ConfirmationParties []*model.ConfirmationParties2 `xml:"ConfPties"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails43 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstruction *model.StandingSettlementInstruction9 `xml:"StgSttlmInstr,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties23 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties23 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the specific transaction.
	CashParties *model.CashParties6 `xml:"CshPties,omitempty"`

	// Provides clearing member information.
	ClearingDetails *model.Clearing3 `xml:"ClrDtls,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.  The amount includes the principal with any commissions and fees or accrued interest.
	SettlementAmount *model.AmountAndDirection28 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts []*model.OtherAmounts16 `xml:"OthrAmts,omitempty"`

	// Other prices than the deal price.
	OtherPrices []*model.OtherPrices1 `xml:"OthrPrics,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties18 `xml:"OthrBizPties,omitempty"`

	// Identifies a transaction that the trading parties are agreeing to repurchase, sell back or return the same or similar securities at a later time.
	// The two leg transaction details defines the closing leg conditions of a two leg transaction. It is also used to define the anticipated closing leg conditions at the time of opening the closed-end transaction.
	//
	//
	TwoLegTransactionDetails *model.TwoLegTransactionDetails1 `xml:"TwoLegTxDtls,omitempty"`

	// Specifies regulatory stipulations that financial institutions must be compliant with in the country, region, and/or area they conduct business.
	RegulatoryStipulations *model.RegulatoryStipulations1 `xml:"RgltryStiptns,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeConfirmationV01) AddIdentification() *model.TransactiontIdentification4 {
	s.Identification = new(model.TransactiontIdentification4)
	return s.Identification
}

func (s *SecuritiesTradeConfirmationV01) AddNumberCount() *model.NumberCount1Choice {
	s.NumberCount = new(model.NumberCount1Choice)
	return s.NumberCount
}

func (s *SecuritiesTradeConfirmationV01) AddReferences() *model.Linkages15 {
	newValue := new(model.Linkages15)
	s.References = append(s.References, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddTradeDetails() *model.Order14 {
	s.TradeDetails = new(model.Order14)
	return s.TradeDetails
}

func (s *SecuritiesTradeConfirmationV01) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeConfirmationV01) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes31 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes31)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeConfirmationV01) AddUnderlyingFinancialInstrument() *model.UnderlyingFinancialInstrument1 {
	newValue := new(model.UnderlyingFinancialInstrument1)
	s.UnderlyingFinancialInstrument = append(s.UnderlyingFinancialInstrument, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddStipulations() *model.FinancialInstrumentStipulations2 {
	s.Stipulations = new(model.FinancialInstrumentStipulations2)
	return s.Stipulations
}

func (s *SecuritiesTradeConfirmationV01) AddConfirmationParties() *model.ConfirmationParties2 {
	newValue := new(model.ConfirmationParties2)
	s.ConfirmationParties = append(s.ConfirmationParties, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddSettlementParameters() *model.SettlementDetails43 {
	s.SettlementParameters = new(model.SettlementDetails43)
	return s.SettlementParameters
}

func (s *SecuritiesTradeConfirmationV01) AddStandingSettlementInstruction() *model.StandingSettlementInstruction9 {
	s.StandingSettlementInstruction = new(model.StandingSettlementInstruction9)
	return s.StandingSettlementInstruction
}

func (s *SecuritiesTradeConfirmationV01) AddDeliveringSettlementParties() *model.SettlementParties23 {
	s.DeliveringSettlementParties = new(model.SettlementParties23)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeConfirmationV01) AddReceivingSettlementParties() *model.SettlementParties23 {
	s.ReceivingSettlementParties = new(model.SettlementParties23)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeConfirmationV01) AddCashParties() *model.CashParties6 {
	s.CashParties = new(model.CashParties6)
	return s.CashParties
}

func (s *SecuritiesTradeConfirmationV01) AddClearingDetails() *model.Clearing3 {
	s.ClearingDetails = new(model.Clearing3)
	return s.ClearingDetails
}

func (s *SecuritiesTradeConfirmationV01) AddSettlementAmount() *model.AmountAndDirection28 {
	s.SettlementAmount = new(model.AmountAndDirection28)
	return s.SettlementAmount
}

func (s *SecuritiesTradeConfirmationV01) AddOtherAmounts() *model.OtherAmounts16 {
	newValue := new(model.OtherAmounts16)
	s.OtherAmounts = append(s.OtherAmounts, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddOtherPrices() *model.OtherPrices1 {
	newValue := new(model.OtherPrices1)
	s.OtherPrices = append(s.OtherPrices, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddOtherBusinessParties() *model.OtherParties18 {
	s.OtherBusinessParties = new(model.OtherParties18)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeConfirmationV01) AddTwoLegTransactionDetails() *model.TwoLegTransactionDetails1 {
	s.TwoLegTransactionDetails = new(model.TwoLegTransactionDetails1)
	return s.TwoLegTransactionDetails
}

func (s *SecuritiesTradeConfirmationV01) AddRegulatoryStipulations() *model.RegulatoryStipulations1 {
	s.RegulatoryStipulations = new(model.RegulatoryStipulations1)
	return s.RegulatoryStipulations
}

func (s *SecuritiesTradeConfirmationV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
