package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03400105 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.034.001.05 Document"`
	Message *SecuritiesFinancingStatusAdviceV05 `xml:"SctiesFincgStsAdvc"`
}

func (d *Document03400105) AddMessage() *SecuritiesFinancingStatusAdviceV05 {
	d.Message = new(SecuritiesFinancingStatusAdviceV05)
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
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingStatusAdviceV05 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications1 `xml:"TxId"`

	// Processing status of the transaction.
	ProcessingStatus *model.ProcessingStatus38Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *model.MatchingStatus20Choice `xml:"MtchgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *model.MatchingStatus20Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *model.SettlementStatus10Choice `xml:"SttlmSts,omitempty"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *model.RepoCallRequestStatus5Choice `xml:"RepoCallReqSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *model.SecuritiesFinancingTransactionDetails24 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingStatusAdviceV05) AddTransactionIdentification() *model.TransactionIdentifications1 {
	s.TransactionIdentification = new(model.TransactionIdentifications1)
	return s.TransactionIdentification
}

func (s *SecuritiesFinancingStatusAdviceV05) AddProcessingStatus() *model.ProcessingStatus38Choice {
	s.ProcessingStatus = new(model.ProcessingStatus38Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesFinancingStatusAdviceV05) AddMatchingStatus() *model.MatchingStatus20Choice {
	s.MatchingStatus = new(model.MatchingStatus20Choice)
	return s.MatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV05) AddInferredMatchingStatus() *model.MatchingStatus20Choice {
	s.InferredMatchingStatus = new(model.MatchingStatus20Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV05) AddSettlementStatus() *model.SettlementStatus10Choice {
	s.SettlementStatus = new(model.SettlementStatus10Choice)
	return s.SettlementStatus
}

func (s *SecuritiesFinancingStatusAdviceV05) AddRepoCallRequestStatus() *model.RepoCallRequestStatus5Choice {
	s.RepoCallRequestStatus = new(model.RepoCallRequestStatus5Choice)
	return s.RepoCallRequestStatus
}

func (s *SecuritiesFinancingStatusAdviceV05) AddTransactionDetails() *model.SecuritiesFinancingTransactionDetails24 {
	s.TransactionDetails = new(model.SecuritiesFinancingTransactionDetails24)
	return s.TransactionDetails
}

func (s *SecuritiesFinancingStatusAdviceV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
