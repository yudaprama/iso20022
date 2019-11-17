package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02500102 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.025.001.02 Document"`
	Message *SecuritiesSettlementTransactionConfirmationV02 `xml:"SctiesSttlmTxConf"`
}

func (d *Document02500102) AddMessage() *SecuritiesSettlementTransactionConfirmationV02 {
	d.Message = new(SecuritiesSettlementTransactionConfirmationV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionConfirmation to an account owner to confirm the partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
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
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionConfirmationV02 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification9 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters2 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails2 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount15 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails24 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties8 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection2 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts9 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties8 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification9 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification9)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddAdditionalParameters() *model.AdditionalParameters2 {
	s.AdditionalParameters = new(model.AdditionalParameters2)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddTradeDetails() *model.SecuritiesTradeDetails2 {
	s.TradeDetails = new(model.SecuritiesTradeDetails2)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes20 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes20)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddQuantityAndAccountDetails() *model.QuantityAndAccount15 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount15)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddSettlementParameters() *model.SettlementDetails24 {
	s.SettlementParameters = new(model.SettlementDetails24)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddDeliveringSettlementParties() *model.SettlementParties11 {
	s.DeliveringSettlementParties = new(model.SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddReceivingSettlementParties() *model.SettlementParties11 {
	s.ReceivingSettlementParties = new(model.SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddCashParties() *model.CashParties8 {
	s.CashParties = new(model.CashParties8)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddSettledAmount() *model.AmountAndDirection2 {
	s.SettledAmount = new(model.AmountAndDirection2)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddOtherAmounts() *model.OtherAmounts9 {
	s.OtherAmounts = new(model.OtherAmounts9)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddOtherBusinessParties() *model.OtherParties8 {
	s.OtherBusinessParties = new(model.OtherParties8)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
