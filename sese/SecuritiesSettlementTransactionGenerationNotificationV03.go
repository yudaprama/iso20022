package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03200103 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.032.001.03 Document"`
	Message *SecuritiesSettlementTransactionGenerationNotificationV03 `xml:"SctiesSttlmTxGnrtnNtfctn"`
}

func (d *Document03200103) AddMessage() *SecuritiesSettlementTransactionGenerationNotificationV03 {
	d.Message = new(SecuritiesSettlementTransactionGenerationNotificationV03)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionGenerationNotification to an account owner to advise the account owner of a securities settlement transaction that has been generated/created by the account servicer for the account owner. The reason for creation can range from the automatic transformation of pending settlement instructions following a corporate event to the recovery of an account owner transactions' database initiated by its account servicer.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// using the relevant elements in the Business Application Header.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionGenerationNotificationV03 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification15 `xml:"TxIdDtls"`

	// Count of the number of transactions linked.
	NumberCounts *model.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*model.Linkages17 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails1 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails []*model.QuantityAndAccount25 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails42 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties8 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *model.AmountAndDirection2 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason []*model.GeneratedReason1 `xml:"GnrtdRsn,omitempty"`

	// Status and reason of the transaction.
	StatusAndReason *model.StatusAndReason12 `xml:"StsAndRsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification15 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification15)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddLinkages() *model.Linkages17 {
	newValue := new(model.Linkages17)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddTradeDetails() *model.SecuritiesTradeDetails1 {
	s.TradeDetails = new(model.SecuritiesTradeDetails1)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddQuantityAndAccountDetails() *model.QuantityAndAccount25 {
	newValue := new(model.QuantityAndAccount25)
	s.QuantityAndAccountDetails = append(s.QuantityAndAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddSettlementParameters() *model.SettlementDetails42 {
	s.SettlementParameters = new(model.SettlementDetails42)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddDeliveringSettlementParties() *model.SettlementParties11 {
	s.DeliveringSettlementParties = new(model.SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddReceivingSettlementParties() *model.SettlementParties11 {
	s.ReceivingSettlementParties = new(model.SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddCashParties() *model.CashParties8 {
	s.CashParties = new(model.CashParties8)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddSettlementAmount() *model.AmountAndDirection2 {
	s.SettlementAmount = new(model.AmountAndDirection2)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddOtherAmounts() *model.OtherAmounts14 {
	s.OtherAmounts = new(model.OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddOtherBusinessParties() *model.OtherParties19 {
	s.OtherBusinessParties = new(model.OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddGeneratedReason() *model.GeneratedReason1 {
	newValue := new(model.GeneratedReason1)
	s.GeneratedReason = append(s.GeneratedReason, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddStatusAndReason() *model.StatusAndReason12 {
	s.StatusAndReason = new(model.StatusAndReason12)
	return s.StatusAndReason
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
