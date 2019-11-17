package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900101 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.029.001.01 Document"`
	Message *SecuritiesSettlementAllegementRemovalAdviceV01 `xml:"SctiesSttlmAllgmtRmvlAdvc"`
}

func (d *Document02900101) AddMessage() *SecuritiesSettlementAllegementRemovalAdviceV01 {
	d.Message = new(SecuritiesSettlementAllegementRemovalAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementAllegementRemovalAdvice to an account owner to acknowledge that a previously sent allegement is no longer outstanding, because the alleged party sent its instruction.
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
type SecuritiesSettlementAllegementRemovalAdviceV01 struct {

	// Information that unambiguously identifies a securities settlement transaction and a SecuritiesSettlementAllegementRemovalAdvice message as known by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Details of the transaction.
	Details *model.TransactionDetails13 `xml:"Dtls"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV01) AddDetails() *model.TransactionDetails13 {
	s.Details = new(model.TransactionDetails13)
	return s.Details
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
