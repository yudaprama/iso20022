package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04400102 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.044.001.02 Document"`
	Message *SecuritiesTradeConfirmationStatusAdviceV02 `xml:"SctiesTradConfStsAdvc"`
}

func (d *Document04400102) AddMessage() *SecuritiesTradeConfirmationStatusAdviceV02 {
	d.Message = new(SecuritiesTradeConfirmationStatusAdviceV02)
	return d.Message
}

// Scope
// This message is sent from Central Matching Utility (CMU) to an executing party or an instructing party to advise the status of the SecuritiesTradeConfirmation message previously sent by the party. The status may be a processing, pending processing, affirmed or disaffirmed, cancel or replacement by an instructing party, a custodian or an affirming party, internal matching, and/or matching status.
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager or of other categories of investors. The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: In central matching the Initiator is the Central Matching Utility.
// Respondent: no response is needed by the recipient of the message.
type SecuritiesTradeConfirmationStatusAdviceV02 struct {

	// Information that unambiguously identifies an SecuritiesTradeConfirmationStatusAdvice message as known by the account owner (or the instructing party acting on its behalf).
	Identification *model.TransactiontIdentification4 `xml:"Id"`

	// Link to another transaction that must be processed after, before or at the same time.
	References []*model.Linkages18 `xml:"Refs"`

	// Provides details on the affitrmation status of a trade.
	AffirmationStatus *model.AffirmationStatus6Choice `xml:"AffirmSts,omitempty"`

	// Provides the processing status of a trade.
	ProcessingStatus *model.ProcessingStatus17Choice `xml:"PrcgSts,omitempty"`

	// Provides details on the matching status of a trade.
	MatchingStatus *model.MatchingStatus23Choice `xml:"MtchgSts,omitempty"`

	// Provides the replacement processing status of a trade.
	ReplacementProcessingStatus *model.ReplacementProcessingStatus7Choice `xml:"RplcmntPrcgSts,omitempty"`

	// Provides details on the cancellation status of a trade.
	CancellationProcessingStatus *model.CancellationProcessingStatus6Choice `xml:"CxlPrcgSts,omitempty"`

	// Details of the trading party.
	PartyTradingDetails *model.Order18 `xml:"PtyTradgDtls,omitempty"`

	// Details of the trading counterparty.
	CounterpartyTradingDetails *model.Order18 `xml:"CtrPtyTradgDtls,omitempty"`

	// Parties used for acting parties that applies either to the whole message or to individual sides.
	ConfirmationParties []*model.ConfirmationParties4 `xml:"ConfPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *model.SettlementParties23 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *model.SettlementParties23 `xml:"RcvgSttlmPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddIdentification() *model.TransactiontIdentification4 {
	s.Identification = new(model.TransactiontIdentification4)
	return s.Identification
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddReferences() *model.Linkages18 {
	newValue := new(model.Linkages18)
	s.References = append(s.References, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddAffirmationStatus() *model.AffirmationStatus6Choice {
	s.AffirmationStatus = new(model.AffirmationStatus6Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddProcessingStatus() *model.ProcessingStatus17Choice {
	s.ProcessingStatus = new(model.ProcessingStatus17Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddMatchingStatus() *model.MatchingStatus23Choice {
	s.MatchingStatus = new(model.MatchingStatus23Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddReplacementProcessingStatus() *model.ReplacementProcessingStatus7Choice {
	s.ReplacementProcessingStatus = new(model.ReplacementProcessingStatus7Choice)
	return s.ReplacementProcessingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddCancellationProcessingStatus() *model.CancellationProcessingStatus6Choice {
	s.CancellationProcessingStatus = new(model.CancellationProcessingStatus6Choice)
	return s.CancellationProcessingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddPartyTradingDetails() *model.Order18 {
	s.PartyTradingDetails = new(model.Order18)
	return s.PartyTradingDetails
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddCounterpartyTradingDetails() *model.Order18 {
	s.CounterpartyTradingDetails = new(model.Order18)
	return s.CounterpartyTradingDetails
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddConfirmationParties() *model.ConfirmationParties4 {
	newValue := new(model.ConfirmationParties4)
	s.ConfirmationParties = append(s.ConfirmationParties, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddDeliveringSettlementParties() *model.SettlementParties23 {
	s.DeliveringSettlementParties = new(model.SettlementParties23)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddReceivingSettlementParties() *model.SettlementParties23 {
	s.ReceivingSettlementParties = new(model.SettlementParties23)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeConfirmationStatusAdviceV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
