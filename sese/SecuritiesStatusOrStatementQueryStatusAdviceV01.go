package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200101 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.01 Document"`
	Message *SecuritiesStatusOrStatementQueryStatusAdviceV01 `xml:"SctiesStsOrStmtQryStsAdvc"`
}

func (d *Document02200101) AddMessage() *SecuritiesStatusOrStatementQueryStatusAdviceV01 {
	d.Message = new(SecuritiesStatusOrStatementQueryStatusAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesStatusOrStatementQueryStatusAdvice to an account owner to advise the status of a status query or statement query previously sent by the account owner.
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
type SecuritiesStatusOrStatementQueryStatusAdviceV01 struct {

	// Information that unambiguously identifies a SecuritiesStatusOrStatementQueryStatusAdvice message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Unambiguous identification of the query as per the account owner.
	QueryReference *model.Identification1 `xml:"QryRef"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	StatusOrStatementRequested *model.StatusOrStatement1Choice `xml:"StsOrStmtReqd,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *model.ProcessingStatus4Choice `xml:"PrcgSts"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddQueryReference() *model.Identification1 {
	s.QueryReference = new(model.Identification1)
	return s.QueryReference
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddAccountOwner() *model.PartyIdentification13Choice {
	s.AccountOwner = new(model.PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddStatusOrStatementRequested() *model.StatusOrStatement1Choice {
	s.StatusOrStatementRequested = new(model.StatusOrStatement1Choice)
	return s.StatusOrStatementRequested
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddProcessingStatus() *model.ProcessingStatus4Choice {
	s.ProcessingStatus = new(model.ProcessingStatus4Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
