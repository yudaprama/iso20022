package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02700101 struct {
	XMLName xml.Name                                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.01 Document"`
	Message *SecuritiesTransactionCancellationRequestStatusAdviceV01 `xml:"SctiesTxCxlReqStsAdvc"`
}

func (d *Document02700101) AddMessage() *SecuritiesTransactionCancellationRequestStatusAdviceV01 {
	d.Message = new(SecuritiesTransactionCancellationRequestStatusAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends an SecuritiesTransactionCancellationRequestStatusAdvice to an account owner to advise the status of a securities transaction cancellation request previously sent by the account owner.
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
type SecuritiesTransactionCancellationRequestStatusAdviceV01 struct {

	// Information that unambiguously identifies a SecuritiesTransactionCancellationRequestStatusAdvice message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	CancellationRequestReference *model.Identification1 `xml:"CxlReqRef"`

	// Unambiguous identification of the transaction as known by the account servicer.
	TransactionIdentification *model.TransactionIdentifications4 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *model.ProcessingStatus2Choice `xml:"PrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails4 `xml:"TxDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddCancellationRequestReference() *model.Identification1 {
	s.CancellationRequestReference = new(model.Identification1)
	return s.CancellationRequestReference
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddTransactionIdentification() *model.TransactionIdentifications4 {
	s.TransactionIdentification = new(model.TransactionIdentifications4)
	return s.TransactionIdentification
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddProcessingStatus() *model.ProcessingStatus2Choice {
	s.ProcessingStatus = new(model.ProcessingStatus2Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddTransactionDetails() *model.TransactionDetails4 {
	s.TransactionDetails = new(model.TransactionDetails4)
	return s.TransactionDetails
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
