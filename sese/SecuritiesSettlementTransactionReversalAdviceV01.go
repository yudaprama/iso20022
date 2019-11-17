package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02600101 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.026.001.01 Document"`
	Message *SecuritiesSettlementTransactionReversalAdviceV01 `xml:"SctiesSttlmTxRvslAdvc"`
}

func (d *Document02600101) AddMessage() *SecuritiesSettlementTransactionReversalAdviceV01 {
	d.Message = new(SecuritiesSettlementTransactionReversalAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionReversalAdvice to an account owner to reverse the confirmation of a partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
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
type SecuritiesSettlementTransactionReversalAdviceV01 struct {

	// Information that unambiguously identifies a SecuritiesSettlementTransactionReversalAdvice message as known by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *model.SettlementTypeAndIdentification1 `xml:"TxIdDtls"`

	// Reference to the unambiguous identification of the confirmation as per the account servicer.
	ConfirmationReference *model.Identification1 `xml:"ConfRef"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters4 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *model.SecuritiesTradeDetails2 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *model.QuantityAndAccount2 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *model.SettlementDetails6 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *model.StandingSettlementInstruction1 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *model.CashParties3 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *model.AmountAndDirection2 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *model.OtherAmounts4 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *model.OtherParties2 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *model.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddTransactionIdentificationDetails() *model.SettlementTypeAndIdentification1 {
	s.TransactionIdentificationDetails = new(model.SettlementTypeAndIdentification1)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddConfirmationReference() *model.Identification1 {
	s.ConfirmationReference = new(model.Identification1)
	return s.ConfirmationReference
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddAdditionalParameters() *model.AdditionalParameters4 {
	s.AdditionalParameters = new(model.AdditionalParameters4)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddTradeDetails() *model.SecuritiesTradeDetails2 {
	s.TradeDetails = new(model.SecuritiesTradeDetails2)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddFinancialInstrumentIdentification() *model.SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(model.SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddQuantityAndAccountDetails() *model.QuantityAndAccount2 {
	s.QuantityAndAccountDetails = new(model.QuantityAndAccount2)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddSettlementParameters() *model.SettlementDetails6 {
	s.SettlementParameters = new(model.SettlementDetails6)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddStandingSettlementInstructionDetails() *model.StandingSettlementInstruction1 {
	s.StandingSettlementInstructionDetails = new(model.StandingSettlementInstruction1)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddDeliveringSettlementParties() *model.SettlementParties5 {
	s.DeliveringSettlementParties = new(model.SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddReceivingSettlementParties() *model.SettlementParties5 {
	s.ReceivingSettlementParties = new(model.SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddCashParties() *model.CashParties3 {
	s.CashParties = new(model.CashParties3)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddSettledAmount() *model.AmountAndDirection2 {
	s.SettledAmount = new(model.AmountAndDirection2)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddOtherAmounts() *model.OtherAmounts4 {
	s.OtherAmounts = new(model.OtherAmounts4)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddOtherBusinessParties() *model.OtherParties2 {
	s.OtherBusinessParties = new(model.OtherParties2)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddAdditionalPhysicalOrRegistrationDetails() *model.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(model.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesSettlementTransactionReversalAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
