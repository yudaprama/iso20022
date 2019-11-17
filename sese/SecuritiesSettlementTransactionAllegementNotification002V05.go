package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02800205 struct {
	XMLName xml.Name                                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.028.002.05 Document"`
	Message *SecuritiesSettlementTransactionAllegementNotification002V05 `xml:"SctiesSttlmTxAllgmtNtfctn"`
}

func (d *Document02800205) AddMessage() *SecuritiesSettlementTransactionAllegementNotification002V05 {
	d.Message = new(SecuritiesSettlementTransactionAllegementNotification002V05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementNotification to an account owner to advise the account owner that a counterparty has alleged an instruction against the account owner's account at the account servicer and that the account servicer could not find the corresponding instruction of the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionAllegementNotification002V05 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.RestrictedFINXMax16Text `xml:"TxId"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *model.SettlementTypeAndAdditionalParameters15 `xml:"SttlmTpAndAddtlParams"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *model.Identification16 `xml:"MktInfrstrctrTxId,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails60 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount49 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails34 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails108 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies cash parties in the framework of a corporate action event.
	CashParties *model.CashParties32 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *model.AmountAndDirection71 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts36 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties31 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*model.RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddSettlementTypeAndAdditionalParameters() *model.SettlementTypeAndAdditionalParameters15 {
	s.SettlementTypeAndAdditionalParameters = new(model.SettlementTypeAndAdditionalParameters15)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddMarketInfrastructureTransactionIdentification() *model.Identification16 {
	s.MarketInfrastructureTransactionIdentification = new(model.Identification16)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddTradeDetails() *model.SecuritiesTradeDetails60 {
	s.TradeDetails = new(model.SecuritiesTradeDetails60)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddFinancialInstrumentIdentification() *model.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddQuantityAndAccountDetails() *model.QuantityAndAccount49 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount49)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails34 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails34)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddSettlementParameters() *model.SettlementDetails108 {
	s.SettlementParameters = new(model.SettlementDetails108)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddDeliveringSettlementParties() *model.SettlementParties44 {
	s.DeliveringSettlementParties = new(model.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddReceivingSettlementParties() *model.SettlementParties44 {
	s.ReceivingSettlementParties = new(model.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddCashParties() *model.CashParties32 {
	s.CashParties = new(model.CashParties32)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddSettlementAmount() *model.AmountAndDirection71 {
	s.SettlementAmount = new(model.AmountAndDirection71)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddOtherAmounts() *model.OtherAmounts36 {
	s.OtherAmounts = new(model.OtherAmounts36)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddOtherBusinessParties() *model.OtherParties31 {
	s.OtherBusinessParties = new(model.OtherParties31)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionAllegementNotification002V05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
