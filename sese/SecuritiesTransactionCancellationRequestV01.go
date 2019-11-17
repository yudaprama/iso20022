package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02000101 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.01 Document"`
	Message *SecuritiesTransactionCancellationRequestV01 `xml:"SctiesTxCxlReq"`
}

func (d *Document02000101) AddMessage() *SecuritiesTransactionCancellationRequestV01 {
	d.Message = new(SecuritiesTransactionCancellationRequestV01)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesTransactionCancellationRequest to an account servicer to request the cancellation of a securities transaction.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct to a central securities depository ot another settlement market infrastructure.
// Usage
// The transaction may be:
// - a securities settlement transaction
// - an intra-position movement
// - a securities financing transaction
// The instruction cannot be:
// - a securities settlement conditions modification (another transaction processing command should be sent to reverse a processing change previously requested).
// - a securities financing modification
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionCancellationRequestV01 struct {

	// Information that unambiguously identifies a SecuritiesTransactionCancellationRequest message as know by the account owner (or the instructing party acting on its behalf).
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Details of the transaction.
	Details *model.TransactionDetails11 `xml:"Dtls"`

	// Party that originated the message, if other than the sender
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesTransactionCancellationRequestV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesTransactionCancellationRequestV01) AddDetails() *model.TransactionDetails11 {
	s.Details = new(model.TransactionDetails11)
	return s.Details
}

func (s *SecuritiesTransactionCancellationRequestV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesTransactionCancellationRequestV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesTransactionCancellationRequestV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
