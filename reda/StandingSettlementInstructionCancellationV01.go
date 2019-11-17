package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05900101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Document"`
	Message *StandingSettlementInstructionCancellationV01 `xml:"StgSttlmInstrCxl"`
}

func (d *Document05900101) AddMessage() *StandingSettlementInstructionCancellationV01 {
	d.Message = new(StandingSettlementInstructionCancellationV01)
	return d.Message
}

// Scope
// An instructing party sends the StandingSettlementInstructionCancellation message to the receiver to cancel a previously sent StandingSettlementInstruction message. The message can also be used to notify a counterparty of the deletion of a standing settlement information.
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
type StandingSettlementInstructionCancellationV01 struct {

	// Reference of this message.
	MessageReferenceIdentification *model.Max35Text `xml:"MsgRefId"`

	// Date on which the SSI is effective.
	EffectiveDateDetails *model.EffectiveDate1 `xml:"FctvDtDtls,omitempty"`

	// Unique and unambiguous master identification known to the sender (or its authorised agent) and receiver (or its authorised agent), below which the SSI will be lodged. This may be an account number or reference to a fund.
	// If no account or reference is available then “NONREF” must be specified.
	AccountIdentification []*model.AccountIdentification26 `xml:"AcctId"`

	// Identifies the market for the standing settlement instruction.
	MarketIdentification *model.MarketIdentificationOrCashPurpose1Choice `xml:"MktId"`

	// Settlement information that helps to identify the standing settlement  instruction for which the cancellation is sent.
	SettlementDetails *model.PartyOrCurrency1Choice `xml:"SttlmDtls"`

	// Reference to a linked message that was previously sent.
	PreviousMessageReference *model.Max35Text `xml:"PrvsMsgRef"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StandingSettlementInstructionCancellationV01) SetMessageReferenceIdentification(value string) {
	s.MessageReferenceIdentification = (*model.Max35Text)(&value)
}

func (s *StandingSettlementInstructionCancellationV01) AddEffectiveDateDetails() *model.EffectiveDate1 {
	s.EffectiveDateDetails = new(model.EffectiveDate1)
	return s.EffectiveDateDetails
}

func (s *StandingSettlementInstructionCancellationV01) AddAccountIdentification() *model.AccountIdentification26 {
	newValue := new(model.AccountIdentification26)
	s.AccountIdentification = append(s.AccountIdentification, newValue)
	return newValue
}

func (s *StandingSettlementInstructionCancellationV01) AddMarketIdentification() *model.MarketIdentificationOrCashPurpose1Choice {
	s.MarketIdentification = new(model.MarketIdentificationOrCashPurpose1Choice)
	return s.MarketIdentification
}

func (s *StandingSettlementInstructionCancellationV01) AddSettlementDetails() *model.PartyOrCurrency1Choice {
	s.SettlementDetails = new(model.PartyOrCurrency1Choice)
	return s.SettlementDetails
}

func (s *StandingSettlementInstructionCancellationV01) SetPreviousMessageReference(value string) {
	s.PreviousMessageReference = (*model.Max35Text)(&value)
}

func (s *StandingSettlementInstructionCancellationV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
