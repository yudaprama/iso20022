package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02600103 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.026.001.03 Document"`
	Message *SecuritiesSettlementTransactionReversalAdviceV03 `xml:"SctiesSttlmTxRvslAdvc"`
}

func (d *Document02600103) AddMessage() *SecuritiesSettlementTransactionReversalAdviceV03 {
	d.Message = new(SecuritiesSettlementTransactionReversalAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionReversalAdvice to an account owner to reverse the confirmation of a partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
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
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionReversalAdviceV03 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification15 `xml:"TxIdDtls"`

	// Reference to the unambiguous identification of the confirmation as per the account servicer.
	ConfirmationReference *model.Identification1 `xml:"ConfRef"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters14 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails2 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount28 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails45 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction5 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties14 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties14 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties9 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection2 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts18 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification15 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification15)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddConfirmationReference() *model.Identification1 {
	s.ConfirmationReference = new(model.Identification1)
	return s.ConfirmationReference
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddAdditionalParameters() *model.AdditionalParameters14 {
	s.AdditionalParameters = new(model.AdditionalParameters14)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddTradeDetails() *model.SecuritiesTradeDetails2 {
	s.TradeDetails = new(model.SecuritiesTradeDetails2)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddQuantityAndAccountDetails() *model.QuantityAndAccount28 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount28)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddSettlementParameters() *model.SettlementDetails45 {
	s.SettlementParameters = new(model.SettlementDetails45)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction5 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction5)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddDeliveringSettlementParties() *model.SettlementParties14 {
	s.DeliveringSettlementParties = new(model.SettlementParties14)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddReceivingSettlementParties() *model.SettlementParties14 {
	s.ReceivingSettlementParties = new(model.SettlementParties14)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddCashParties() *model.CashParties9 {
	s.CashParties = new(model.CashParties9)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddSettledAmount() *model.AmountAndDirection2 {
	s.SettledAmount = new(model.AmountAndDirection2)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddOtherAmounts() *model.OtherAmounts18 {
	s.OtherAmounts = new(model.OtherAmounts18)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddOtherBusinessParties() *model.OtherParties19 {
	s.OtherBusinessParties = new(model.OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
