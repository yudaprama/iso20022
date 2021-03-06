package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02500107 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.025.001.07 Document"`
	Message *SecuritiesSettlementTransactionConfirmationV07 `xml:"SctiesSttlmTxConf"`
}

func (d *Document02500107) AddMessage() *SecuritiesSettlementTransactionConfirmationV07 {
	d.Message = new(SecuritiesSettlementTransactionConfirmationV07)
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
type SecuritiesSettlementTransactionConfirmationV07 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification19 `xml:"TxIdDtls"`

	// Link to another transaction - provided for information only.
	Linkages *model.Linkages41 `xml:"Lnkgs,omitempty"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters29 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails53 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount41 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails128 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties26 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection46 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts30 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification19 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification19)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddLinkages() *model.Linkages41 {
	s.Linkages = new(model.Linkages41)
	return s.Linkages
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddAdditionalParameters() *model.AdditionalParameters29 {
	s.AdditionalParameters = new(model.AdditionalParameters29)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddTradeDetails() *model.SecuritiesTradeDetails53 {
	s.TradeDetails = new(model.SecuritiesTradeDetails53)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddFinancialInstrumentIdentification() *model.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddQuantityAndAccountDetails() *model.QuantityAndAccount41 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount41)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddSettlementParameters() *model.SettlementDetails128 {
	s.SettlementParameters = new(model.SettlementDetails128)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddDeliveringSettlementParties() *model.SettlementParties36 {
	s.DeliveringSettlementParties = new(model.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddReceivingSettlementParties() *model.SettlementParties36 {
	s.ReceivingSettlementParties = new(model.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddCashParties() *model.CashParties26 {
	s.CashParties = new(model.CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddSettledAmount() *model.AmountAndDirection46 {
	s.SettledAmount = new(model.AmountAndDirection46)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddOtherAmounts() *model.OtherAmounts30 {
	s.OtherAmounts = new(model.OtherAmounts30)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddOtherBusinessParties() *model.OtherParties27 {
	s.OtherBusinessParties = new(model.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
