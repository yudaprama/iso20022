package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02400101 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.024.001.01 Document"`
	Message *SecuritiesSettlementTransactionStatusAdviceV01 `xml:"SctiesSttlmTxStsAdvc"`
}

func (d *Document02400101) AddMessage() *SecuritiesSettlementTransactionStatusAdviceV01 {
	d.Message = new(SecuritiesSettlementTransactionStatusAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionStatusAdvice to an account owner to advise the status of a securities settlement transaction instruction previously sent by the account owner or the status of a settlement transaction existing in the books of the servicer for the account of the owner. The status may be a processing, pending processing, internal matching, matching and/or settlement status.
// The status advice may be sent as a response to the request of the account owner or not.
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
type SecuritiesSettlementTransactionStatusAdviceV01 struct {

	// Information that unambiguously identifies a SecuritiesSettlementTransactionStatusAdvice message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications2 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *model.ProcessingStatus1Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *model.MatchingStatus2Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *model.MatchingStatus2Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *model.SettlementStatus2Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails5 `xml:"TxDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddTransactionIdentification() *model.TransactionIdentifications2 {
	s.TransactionIdentification = new(model.TransactionIdentifications2)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddProcessingStatus() *model.ProcessingStatus1Choice {
	s.ProcessingStatus = new(model.ProcessingStatus1Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddInferredMatchingStatus() *model.MatchingStatus2Choice {
	s.InferredMatchingStatus = new(model.MatchingStatus2Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddMatchingStatus() *model.MatchingStatus2Choice {
	s.MatchingStatus = new(model.MatchingStatus2Choice)
	return s.MatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddSettlementStatus() *model.SettlementStatus2Choice {
	s.SettlementStatus = new(model.SettlementStatus2Choice)
	return s.SettlementStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddTransactionDetails() *model.TransactionDetails5 {
	s.TransactionDetails = new(model.TransactionDetails5)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesSettlementTransactionStatusAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
