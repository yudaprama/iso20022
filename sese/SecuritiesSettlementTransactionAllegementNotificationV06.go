package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02800106 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.028.001.06 Document"`
	Message *SecuritiesSettlementTransactionAllegementNotificationV06 `xml:"SctiesSttlmTxAllgmtNtfctn"`
}

func (d *Document02800106) AddMessage() *SecuritiesSettlementTransactionAllegementNotificationV06 {
	d.Message = new(SecuritiesSettlementTransactionAllegementNotificationV06)
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
type SecuritiesSettlementTransactionAllegementNotificationV06 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *model.SettlementTypeAndAdditionalParameters12 `xml:"SttlmTpAndAddtlParams"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *model.Identification14 `xml:"MktInfrstrctrTxId,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails54 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount42 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails29 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails125 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies cash parties in the framework of a corporate action event.
	CashParties *model.CashParties25 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *model.AmountAndDirection48 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts32 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties28 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*model.Max35Text)(&value)
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddSettlementTypeAndAdditionalParameters() *model.SettlementTypeAndAdditionalParameters12 {
	s.SettlementTypeAndAdditionalParameters = new(model.SettlementTypeAndAdditionalParameters12)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddMarketInfrastructureTransactionIdentification() *model.Identification14 {
	s.MarketInfrastructureTransactionIdentification = new(model.Identification14)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddTradeDetails() *model.SecuritiesTradeDetails54 {
	s.TradeDetails = new(model.SecuritiesTradeDetails54)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddFinancialInstrumentIdentification() *model.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddQuantityAndAccountDetails() *model.QuantityAndAccount42 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount42)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails29 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails29)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddSettlementParameters() *model.SettlementDetails125 {
	s.SettlementParameters = new(model.SettlementDetails125)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddDeliveringSettlementParties() *model.SettlementParties36 {
	s.DeliveringSettlementParties = new(model.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddReceivingSettlementParties() *model.SettlementParties36 {
	s.ReceivingSettlementParties = new(model.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddCashParties() *model.CashParties25 {
	s.CashParties = new(model.CashParties25)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddSettlementAmount() *model.AmountAndDirection48 {
	s.SettlementAmount = new(model.AmountAndDirection48)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddOtherAmounts() *model.OtherAmounts32 {
	s.OtherAmounts = new(model.OtherAmounts32)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddOtherBusinessParties() *model.OtherParties28 {
	s.OtherBusinessParties = new(model.OtherParties28)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
