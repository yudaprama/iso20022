package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03400107 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.034.001.07 Document"`
	Message *SecuritiesFinancingStatusAdviceV07 `xml:"SctiesFincgStsAdvc"`
}

func (d *Document03400107) AddMessage() *SecuritiesFinancingStatusAdviceV07 {
	d.Message = new(SecuritiesFinancingStatusAdviceV07)
	return d.Message
}

// Scope
// An securities financing transaction account servicer sends a SecuritiesFinancingStatusAdvice to an account owner to advise the status of a securities financing transaction previously instructed by the account owner.
// The status advice may be sent as a response to the request of the account owner or not.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesFinancingStatusAdviceV07 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications32 `xml:"TxId"`

	// Processing status of the transaction.
	ProcessingStatus *model.ProcessingStatus51Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *model.MatchingStatus26Choice `xml:"MtchgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *model.MatchingStatus26Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *model.SettlementStatus18Choice `xml:"SttlmSts,omitempty"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *model.RepoCallRequestStatus7Choice `xml:"RepoCallReqSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *model.SecuritiesFinancingTransactionDetails35 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingStatusAdviceV07) AddTransactionIdentification() *model.TransactionIdentifications32 {
	s.TransactionIdentification = new(model.TransactionIdentifications32)
	return s.TransactionIdentification
}

func (s *SecuritiesFinancingStatusAdviceV07) AddProcessingStatus() *model.ProcessingStatus51Choice {
	s.ProcessingStatus = new(model.ProcessingStatus51Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesFinancingStatusAdviceV07) AddMatchingStatus() *model.MatchingStatus26Choice {
	s.MatchingStatus = new(model.MatchingStatus26Choice)
	return s.MatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV07) AddInferredMatchingStatus() *model.MatchingStatus26Choice {
	s.InferredMatchingStatus = new(model.MatchingStatus26Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV07) AddSettlementStatus() *model.SettlementStatus18Choice {
	s.SettlementStatus = new(model.SettlementStatus18Choice)
	return s.SettlementStatus
}

func (s *SecuritiesFinancingStatusAdviceV07) AddRepoCallRequestStatus() *model.RepoCallRequestStatus7Choice {
	s.RepoCallRequestStatus = new(model.RepoCallRequestStatus7Choice)
	return s.RepoCallRequestStatus
}

func (s *SecuritiesFinancingStatusAdviceV07) AddTransactionDetails() *model.SecuritiesFinancingTransactionDetails35 {
	s.TransactionDetails = new(model.SecuritiesFinancingTransactionDetails35)
	return s.TransactionDetails
}

func (s *SecuritiesFinancingStatusAdviceV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
