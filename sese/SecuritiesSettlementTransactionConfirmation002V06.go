package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02500206 struct {
	XMLName xml.Name                                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.025.002.06 Document"`
	Message *SecuritiesSettlementTransactionConfirmation002V06 `xml:"SctiesSttlmTxConf"`
}

func (d *Document02500206) AddMessage() *SecuritiesSettlementTransactionConfirmation002V06 {
	d.Message = new(SecuritiesSettlementTransactionConfirmation002V06)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionConfirmation to an account owner to confirm the partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionConfirmation002V06 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification24 `xml:"TxIdDtls"`

	// Link to another transaction - provided for information only.
	Linkages *model.Linkages50 `xml:"Lnkgs,omitempty"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters27 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails62 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount51 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails110 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties30 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection60 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts38 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification24 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification24)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddLinkages() *model.Linkages50 {
	s.Linkages = new(model.Linkages50)
	return s.Linkages
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddAdditionalParameters() *model.AdditionalParameters27 {
	s.AdditionalParameters = new(model.AdditionalParameters27)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddTradeDetails() *model.SecuritiesTradeDetails62 {
	s.TradeDetails = new(model.SecuritiesTradeDetails62)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddFinancialInstrumentIdentification() *model.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddQuantityAndAccountDetails() *model.QuantityAndAccount51 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount51)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddSettlementParameters() *model.SettlementDetails110 {
	s.SettlementParameters = new(model.SettlementDetails110)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddDeliveringSettlementParties() *model.SettlementParties44 {
	s.DeliveringSettlementParties = new(model.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddReceivingSettlementParties() *model.SettlementParties44 {
	s.ReceivingSettlementParties = new(model.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddCashParties() *model.CashParties30 {
	s.CashParties = new(model.CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddSettledAmount() *model.AmountAndDirection60 {
	s.SettledAmount = new(model.AmountAndDirection60)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddOtherAmounts() *model.OtherAmounts38 {
	s.OtherAmounts = new(model.OtherAmounts38)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddOtherBusinessParties() *model.OtherParties29 {
	s.OtherBusinessParties = new(model.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionConfirmation002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
