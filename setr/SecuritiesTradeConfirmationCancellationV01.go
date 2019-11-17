package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900101 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.029.001.01 Document"`
	Message *SecuritiesTradeConfirmationCancellationV01 `xml:"SctiesTradConfCxl"`
}

func (d *Document02900101) AddMessage() *SecuritiesTradeConfirmationCancellationV01 {
	d.Message = new(SecuritiesTradeConfirmationCancellationV01)
	return d.Message
}

// Scope
// Sent by an executing party to an instructing party directly or through Central Matching Utility (CMU) to cancel the referenced SecuritiesTradeConfirmation message that was previously sent.
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager or of other categories of investors.
// The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// It may also be used to cancel the trade confirmation previously sent from an executing party or an instructing party to a custodian or an affirming party directly or through CMU.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: Both in local and central matching, the Initiator may be either the Executing Party or Instructing Party.
// Respondent: Instructing party, a custodian or an affirming party optionally responds with SecuritiesTradeConfirmationResponse (accept or reject) message.
type SecuritiesTradeConfirmationCancellationV01 struct {

	// Information that unambiguously identifies an SecuritiesTradeConfirmationCancellation message as known by the account owner (or the instructing party acting on its behalf).
	Identification *model.TransactiontIdentification4 `xml:"Id"`

	// Link to another transaction that must be processed after, before or at the same time.
	References []*model.Linkages15 `xml:"Refs,omitempty"`

	// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
	OtherBusinessParties *model.OtherParties18 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeConfirmationCancellationV01) AddIdentification() *model.TransactiontIdentification4 {
	s.Identification = new(model.TransactiontIdentification4)
	return s.Identification
}

func (s *SecuritiesTradeConfirmationCancellationV01) AddReferences() *model.Linkages15 {
	newValue := new(model.Linkages15)
	s.References = append(s.References, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationCancellationV01) AddOtherBusinessParties() *model.OtherParties18 {
	s.OtherBusinessParties = new(model.OtherParties18)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeConfirmationCancellationV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
