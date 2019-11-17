package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02800101 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.028.001.01 Document"`
	Message *SecuritiesSettlementTransactionAllegementNotificationV01 `xml:"SctiesSttlmTxAllgmtNtfctn"`
}

func (d *Document02800101) AddMessage() *SecuritiesSettlementTransactionAllegementNotificationV01 {
	d.Message = new(SecuritiesSettlementTransactionAllegementNotificationV01)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementNotification to an account owner to advise the account owner that a counterparty has alleged an instruction against the account owner's account at the account servicer and that the account servicer could not find the corresponding instruction of the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionAllegementNotificationV01 struct {

	// Information that unambiguously identifies a SecuritiesSettlementTransaction and a SecuritiesSettlementTransactionAllegementNotification message as known by the account servicer.
	Identification *model.TransactionAndDocumentIdentification1 `xml:"Id"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *model.SettlementTypeAndAdditionalParameters2 `xml:"SttlmTpAndAddtlParams"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *model.Identification1 `xml:"MktInfrstrctrTxId,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails8 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount3 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *model.SecuritiesFinancingTransactionDetails1 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails5 `xml:"SttlmParams"`

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
	OtherBusinessParties *model.OtherParties3 `xml:"OthrBizPties,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddIdentification() *model.TransactionAndDocumentIdentification1 {
	s.Identification = new(model.TransactionAndDocumentIdentification1)
	return s.Identification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddSettlementTypeAndAdditionalParameters() *model.SettlementTypeAndAdditionalParameters2 {
	s.SettlementTypeAndAdditionalParameters = new(model.SettlementTypeAndAdditionalParameters2)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddMarketInfrastructureTransactionIdentification() *model.Identification1 {
	s.MarketInfrastructureTransactionIdentification = new(model.Identification1)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddTradeDetails() *model.SecuritiesTradeDetails8 {
	s.TradeDetails = new(model.SecuritiesTradeDetails8)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddFinancialInstrumentIdentification() *model.SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddQuantityAndAccountDetails() *model.QuantityAndAccount3 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount3)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddSecuritiesFinancingDetails() *model.SecuritiesFinancingTransactionDetails1 {
	s.SecuritiesFinancingDetails = new(model.SecuritiesFinancingTransactionDetails1)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddSettlementParameters() *model.SettlementDetails5 {
	s.SettlementParameters = new(model.SettlementDetails5)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddDeliveringSettlementParties() *model.SettlementParties5 {
	s.DeliveringSettlementParties = new(model.SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddReceivingSettlementParties() *model.SettlementParties5 {
	s.ReceivingSettlementParties = new(model.SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddCashParties() *model.CashParties3 {
	s.CashParties = new(model.CashParties3)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddSettlementAmount() *model.AmountAndDirection2 {
	s.SettlementAmount = new(model.AmountAndDirection2)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddOtherAmounts() *model.OtherAmounts3 {
	s.OtherAmounts = new(model.OtherAmounts3)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddOtherBusinessParties() *model.OtherParties3 {
	s.OtherBusinessParties = new(model.OtherParties3)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
