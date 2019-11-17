package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03200106 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.032.001.06 Document"`
	Message *SecuritiesSettlementTransactionGenerationNotificationV06 `xml:"SctiesSttlmTxGnrtnNtfctn"`
}

func (d *Document03200106) AddMessage() *SecuritiesSettlementTransactionGenerationNotificationV06 {
	d.Message = new(SecuritiesSettlementTransactionGenerationNotificationV06)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionGenerationNotification to an account owner to advise the account owner of a securities settlement transaction that has been generated/created by the account servicer for the account owner. The reason for creation can range from the automatic transformation of pending settlement instructions following a corporate event to the recovery of an account owner transactions' database initiated by its account servicer.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionGenerationNotificationV06 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification19 `xml:"TxIdDtls"`

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
	QuantityAndAccountDetails []*model.QuantityAndAccount39 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails101 `xml:"SttlmParams"`

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

	// Specifies the reason why the transaction was generated.
	GeneratedReason []*model.GeneratedReason5 `xml:"GnrtdRsn,omitempty"`

	// Status and reason of the transaction.
	StatusAndReason *model.StatusAndReason28 `xml:"StsAndRsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification19 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification19)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddLinkages() *model.Linkages37 {
	newValue := new(model.Linkages37)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddTradeDetails() *model.SecuritiesTradeDetails51 {
	s.TradeDetails = new(model.SecuritiesTradeDetails51)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddFinancialInstrumentIdentification() *model.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddQuantityAndAccountDetails() *model.QuantityAndAccount39 {
	newValue := new(model.QuantityAndAccount39)
	s.QuantityAndAccountDetails = append(s.QuantityAndAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddSettlementParameters() *model.SettlementDetails101 {
	s.SettlementParameters = new(model.SettlementDetails101)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddDeliveringSettlementParties() *model.SettlementParties36 {
	s.DeliveringSettlementParties = new(model.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddReceivingSettlementParties() *model.SettlementParties36 {
	s.ReceivingSettlementParties = new(model.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddCashParties() *model.CashParties26 {
	s.CashParties = new(model.CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddSettlementAmount() *model.AmountAndDirection46 {
	s.SettlementAmount = new(model.AmountAndDirection46)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddOtherAmounts() *model.OtherAmounts28 {
	s.OtherAmounts = new(model.OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddOtherBusinessParties() *model.OtherParties27 {
	s.OtherBusinessParties = new(model.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddGeneratedReason() *model.GeneratedReason5 {
	newValue := new(model.GeneratedReason5)
	s.GeneratedReason = append(s.GeneratedReason, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddStatusAndReason() *model.StatusAndReason28 {
	s.StatusAndReason = new(model.StatusAndReason28)
	return s.StatusAndReason
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
