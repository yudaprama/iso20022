package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03000101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.030.001.01 Document"`
	Message *SecuritiesTradeConfirmationResponseV01 `xml:"SctiesTradConfRspn"`
}

func (d *Document03000101) AddMessage() *SecuritiesTradeConfirmationResponseV01 {
	d.Message = new(SecuritiesTradeConfirmationResponseV01)
	return d.Message
}

// Scope
// Sent by an instructing party, a custodian or an affirming party to an executing party (local matching) or to Central Matching Utility (CMU) to affirm (accept) or disaffirm (reject) (central matching) the SecuritiesTradeConfirmation message. If accepting the SecuritiesTradeConfirmation message, then the trade is ready for settlement processing. If rejecting the SecuritiesTradeConfirmation message, then the trade is not ready for settlement.
// The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager or of other categories of investors.
// The custodian or an affirming party is typically the custodian, trustee, financial institution, intermediary system/vendor communicating on behalf of them, or their agent.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: Both in local and central matching, the Initiator may be the Instructing Party, Custodian or Affirming party.
// Respondent: Executing party does not need to respond if an affirmation. Executing party may respond with modification or cancellation of the rejected SecuritiesTradeConfirmation message.
type SecuritiesTradeConfirmationResponseV01 struct {

	// Information that unambiguously identifies an SecuritiesTradeConfirmationResponse message as known by the account owner (or the instructing party acting on its behalf).
	Identification *model.TransactiontIdentification4 `xml:"Id"`

	// Link to another transaction that must be processed after, before or at the same time.
	References []*model.Linkages15 `xml:"Refs"`

	// Provides details on the processing status of the trade.
	Status *model.StatusAndReason10 `xml:"Sts"`

	// Provides clearing member information.
	ClearingDetails *model.Clearing3 `xml:"ClrDtls,omitempty"`

	// Parties involved in the confirmation of the details of a trade.
	ConfirmationParties []*model.ConfirmationParties3 `xml:"ConfPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeConfirmationResponseV01) AddIdentification() *model.TransactiontIdentification4 {
	s.Identification = new(model.TransactiontIdentification4)
	return s.Identification
}

func (s *SecuritiesTradeConfirmationResponseV01) AddReferences() *model.Linkages15 {
	newValue := new(model.Linkages15)
	s.References = append(s.References, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationResponseV01) AddStatus() *model.StatusAndReason10 {
	s.Status = new(model.StatusAndReason10)
	return s.Status
}

func (s *SecuritiesTradeConfirmationResponseV01) AddClearingDetails() *model.Clearing3 {
	s.ClearingDetails = new(model.Clearing3)
	return s.ClearingDetails
}

func (s *SecuritiesTradeConfirmationResponseV01) AddConfirmationParties() *model.ConfirmationParties3 {
	newValue := new(model.ConfirmationParties3)
	s.ConfirmationParties = append(s.ConfirmationParties, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationResponseV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
