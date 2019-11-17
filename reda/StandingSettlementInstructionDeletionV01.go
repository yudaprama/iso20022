package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05700101 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.057.001.01 Document"`
	Message *StandingSettlementInstructionDeletionV01 `xml:"StgSttlmInstrDeltn"`
}

func (d *Document05700101) AddMessage() *StandingSettlementInstructionDeletionV01 {
	d.Message = new(StandingSettlementInstructionDeletionV01)
	return d.Message
}

// Scope
// An instructing party sends the StandingSettlementInstructionDeletion message to the receiver to delete a previously sent StandingSettlementInstruction message. The message can also be used to notify a counterparty of the deletion of a standing settlement information.
//
// Usage
// The instructing party (initiator) is:
// • An account servicer, for example, a global custodian or prime broker
// • A counterparty in a transaction, for example:
// 	- an investment manager (executing broker),
// 	- a global custodian (executing broker, prime broker)
// • A vendor's application communicating on behalf of an account servicer or counterparty
// The receiver is:
// • An account owner, for example, an investment manager, hedge fund administrator or a party to which SSI operations have been outsourced
// • A counterparty, for example:
// 	- an investment manager (executing broker)
// 	- a global custodian (executing broker, prime broker)
// • A vendor's application communicating on behalf of the account owner or counterparty.
type StandingSettlementInstructionDeletionV01 struct {

	// Reference of this message.
	MessageReferenceIdentification *model.Max35Text `xml:"MsgRefId"`

	// Date on which the SSI is effective.
	EffectiveDateDetails *model.EffectiveDate1 `xml:"FctvDtDtls,omitempty"`

	// Unique and unambiguous master identification known to the sender (or its authorised agent) and receiver (or its authorised agent), below which the SSI will be lodged. This may be an account number or reference to a fund.
	// If no account or reference is available then “NONREF” must be specified.
	AccountIdentification []*model.AccountIdentification26 `xml:"AcctId"`

	// Identifies the market for the standing settlement instruction.
	MarketIdentification *model.MarketIdentificationOrCashPurpose1Choice `xml:"MktId"`

	// Settlement information that helps to identify the standing settlement  instruction for which the deletion is sent.
	SettlementDetails *model.PartyOrCurrency1Choice `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StandingSettlementInstructionDeletionV01) SetMessageReferenceIdentification(value string) {
	s.MessageReferenceIdentification = (*model.Max35Text)(&value)
}

func (s *StandingSettlementInstructionDeletionV01) AddEffectiveDateDetails() *model.EffectiveDate1 {
	s.EffectiveDateDetails = new(model.EffectiveDate1)
	return s.EffectiveDateDetails
}

func (s *StandingSettlementInstructionDeletionV01) AddAccountIdentification() *model.AccountIdentification26 {
	newValue := new(model.AccountIdentification26)
	s.AccountIdentification = append(s.AccountIdentification, newValue)
	return newValue
}

func (s *StandingSettlementInstructionDeletionV01) AddMarketIdentification() *model.MarketIdentificationOrCashPurpose1Choice {
	s.MarketIdentification = new(model.MarketIdentificationOrCashPurpose1Choice)
	return s.MarketIdentification
}

func (s *StandingSettlementInstructionDeletionV01) AddSettlementDetails() *model.PartyOrCurrency1Choice {
	s.SettlementDetails = new(model.PartyOrCurrency1Choice)
	return s.SettlementDetails
}

func (s *StandingSettlementInstructionDeletionV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
