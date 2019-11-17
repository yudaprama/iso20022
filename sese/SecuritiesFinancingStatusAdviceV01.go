package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03400101 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.034.001.01 Document"`
	Message *SecuritiesFinancingStatusAdviceV01 `xml:"SctiesFincgStsAdvc"`
}

func (d *Document03400101) AddMessage() *SecuritiesFinancingStatusAdviceV01 {
	d.Message = new(SecuritiesFinancingStatusAdviceV01)
	return d.Message
}

// Scope
// An securities financing transaction account servicer sends a SecuritiesFinancingStatusAdvice to an account owner to advise the status of a securities financing transaction previously instructed by the account owner.
// The status advice may be sent as a response to the request of the account owner or not.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingStatusAdviceV01 struct {

	// Information that unambiguously identifies a SecuritiesFinancingStatusAdvice message as known by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications1 `xml:"TxId"`

	// Processing status of the transaction.
	ProcessingStatus *model.ProcessingStatus3Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *model.MatchingStatus3Choice `xml:"MtchgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *model.MatchingStatus3Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *model.SettlementStatus2Choice `xml:"SttlmSts,omitempty"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *model.RepoCallRequestStatus2Choice `xml:"RepoCallReqSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *model.SecuritiesFinancingTransactionDetails2 `xml:"TxDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesFinancingStatusAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesFinancingStatusAdviceV01) AddTransactionIdentification() *model.TransactionIdentifications1 {
	s.TransactionIdentification = new(model.TransactionIdentifications1)
	return s.TransactionIdentification
}

func (s *SecuritiesFinancingStatusAdviceV01) AddProcessingStatus() *model.ProcessingStatus3Choice {
	s.ProcessingStatus = new(model.ProcessingStatus3Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesFinancingStatusAdviceV01) AddMatchingStatus() *model.MatchingStatus3Choice {
	s.MatchingStatus = new(model.MatchingStatus3Choice)
	return s.MatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV01) AddInferredMatchingStatus() *model.MatchingStatus3Choice {
	s.InferredMatchingStatus = new(model.MatchingStatus3Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV01) AddSettlementStatus() *model.SettlementStatus2Choice {
	s.SettlementStatus = new(model.SettlementStatus2Choice)
	return s.SettlementStatus
}

func (s *SecuritiesFinancingStatusAdviceV01) AddRepoCallRequestStatus() *model.RepoCallRequestStatus2Choice {
	s.RepoCallRequestStatus = new(model.RepoCallRequestStatus2Choice)
	return s.RepoCallRequestStatus
}

func (s *SecuritiesFinancingStatusAdviceV01) AddTransactionDetails() *model.SecuritiesFinancingTransactionDetails2 {
	s.TransactionDetails = new(model.SecuritiesFinancingTransactionDetails2)
	return s.TransactionDetails
}

func (s *SecuritiesFinancingStatusAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesFinancingStatusAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesFinancingStatusAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
