package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05800101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.058.001.01 Document"`
	Message *StandingSettlementInstructionStatusAdviceV01 `xml:"StgSttlmInstrStsAdvc"`
}

func (d *Document05800101) AddMessage() *StandingSettlementInstructionStatusAdviceV01 {
	d.Message = new(StandingSettlementInstructionStatusAdviceV01)
	return d.Message
}

// Scope
// The receiver of a StandingSettlementInstruction message sends the StandingSettlementInstructionStatusAdvice message to the instructing party (sender of the StandingSettlementInstruction message) to provide the status of a previously received StandingSettlementInstruction, StandingSettlementInstructionCancellation or StandingSettlementInstructionDeletion message.
//
// Usage
// The StandingSettlementInstructionStatusAdvice message is used to report one of the following statuses:
// -	a received status, or,
// -	an accepted status, or,
// -	a rejected status, or,
// -	a pending processing status, or,
// -	a proprietary status.
type StandingSettlementInstructionStatusAdviceV01 struct {

	// Date on which the SSI is effective.
	EffectiveDateDetails *model.EffectiveDate1 `xml:"FctvDtDtls,omitempty"`

	// Unique and unambiguous master identification known to the sender (or its authorised agent) and receiver (or its authorised agent), below which the SSI will be lodged. This may be an account number or reference to a fund.
	// If no account or reference is available then “NONREF” must be specified.
	AccountIdentification []*model.AccountIdentification26 `xml:"AcctId"`

	// Identifies the market for the standing settlement instruction.
	MarketIdentification *model.MarketIdentificationOrCashPurpose1Choice `xml:"MktId"`

	// Settlement information that helps to identify the standing settlement  instruction, cancellation or deletion for which the status is sent.
	SettlementDetails *model.PartyOrCurrency1Choice `xml:"SttlmDtls"`

	// Reference to a linked message that was previously received.
	RelatedMessageReference *model.Max35Text `xml:"RltdMsgRef"`

	// Status of the standing settlement instruction, deletion or cancellation.
	ProcessingStatus *model.ProcessingStatus43Choice `xml:"PrcgSts"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddEffectiveDateDetails() *model.EffectiveDate1 {
	s.EffectiveDateDetails = new(model.EffectiveDate1)
	return s.EffectiveDateDetails
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddAccountIdentification() *model.AccountIdentification26 {
	newValue := new(model.AccountIdentification26)
	s.AccountIdentification = append(s.AccountIdentification, newValue)
	return newValue
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddMarketIdentification() *model.MarketIdentificationOrCashPurpose1Choice {
	s.MarketIdentification = new(model.MarketIdentificationOrCashPurpose1Choice)
	return s.MarketIdentification
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddSettlementDetails() *model.PartyOrCurrency1Choice {
	s.SettlementDetails = new(model.PartyOrCurrency1Choice)
	return s.SettlementDetails
}

func (s *StandingSettlementInstructionStatusAdviceV01) SetRelatedMessageReference(value string) {
	s.RelatedMessageReference = (*model.Max35Text)(&value)
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddProcessingStatus() *model.ProcessingStatus43Choice {
	s.ProcessingStatus = new(model.ProcessingStatus43Choice)
	return s.ProcessingStatus
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
