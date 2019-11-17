package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03200206 struct {
	XMLName xml.Name                                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.032.002.06 Document"`
	Message *SecuritiesSettlementTransactionGenerationNotification002V06 `xml:"SctiesSttlmTxGnrtnNtfctn"`
}

func (d *Document03200206) AddMessage() *SecuritiesSettlementTransactionGenerationNotification002V06 {
	d.Message = new(SecuritiesSettlementTransactionGenerationNotification002V06)
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
type SecuritiesSettlementTransactionGenerationNotification002V06 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification24 `xml:"TxIdDtls"`

	// Count of the number of transactions linked.
	NumberCounts *model.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*model.Linkages43 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails63 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails []*model.QuantityAndAccount47 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails111 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *model.AmountAndDirection60 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason []*model.GeneratedReason6 `xml:"GnrtdRsn,omitempty"`

	// Status and reason of the transaction.
	StatusAndReason *model.StatusAndReason29 `xml:"StsAndRsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification24 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification24)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddNumberCounts() *model.NumberCount1Choice {
	s.NumberCounts = new(model.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddLinkages() *model.Linkages43 {
	newValue := new(model.Linkages43)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddTradeDetails() *model.SecuritiesTradeDetails63 {
	s.TradeDetails = new(model.SecuritiesTradeDetails63)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddFinancialInstrumentIdentification() *model.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddQuantityAndAccountDetails() *model.QuantityAndAccount47 {
	newValue := new(model.QuantityAndAccount47)
	s.QuantityAndAccountDetails = append(s.QuantityAndAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddSettlementParameters() *model.SettlementDetails111 {
	s.SettlementParameters = new(model.SettlementDetails111)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddDeliveringSettlementParties() *model.SettlementParties44 {
	s.DeliveringSettlementParties = new(model.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddReceivingSettlementParties() *model.SettlementParties44 {
	s.ReceivingSettlementParties = new(model.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddCashParties() *model.CashParties30 {
	s.CashParties = new(model.CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddSettlementAmount() *model.AmountAndDirection60 {
	s.SettlementAmount = new(model.AmountAndDirection60)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddOtherAmounts() *model.OtherAmounts35 {
	s.OtherAmounts = new(model.OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddOtherBusinessParties() *model.OtherParties29 {
	s.OtherBusinessParties = new(model.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddGeneratedReason() *model.GeneratedReason6 {
	newValue := new(model.GeneratedReason6)
	s.GeneratedReason = append(s.GeneratedReason, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddStatusAndReason() *model.StatusAndReason29 {
	s.StatusAndReason = new(model.StatusAndReason29)
	return s.StatusAndReason
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
