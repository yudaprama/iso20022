package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200102 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.02 Document"`
	Message *SecuritiesStatusOrStatementQueryStatusAdviceV02 `xml:"SctiesStsOrStmtQryStsAdvc"`
}

func (d *Document02200102) AddMessage() *SecuritiesStatusOrStatementQueryStatusAdviceV02 {
	d.Message = new(SecuritiesStatusOrStatementQueryStatusAdviceV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesStatusOrStatementQueryStatusAdvice to an account owner to advise the status of a status query or statement query previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// using the relevant elements in the Business Application Header.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesStatusOrStatementQueryStatusAdviceV02 struct {

	// Unambiguous identification of the query as per the account owner.
	QueryReference *model.Identification1 `xml:"QryRef"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	StatusOrStatementRequested *model.StatusOrStatement3Choice `xml:"StsOrStmtReqd,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *model.ProcessingStatus4Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV02) AddQueryReference() *model.Identification1 {
	s.QueryReference = new(model.Identification1)
	return s.QueryReference
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV02) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV02) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV02) AddStatusOrStatementRequested() *model.StatusOrStatement3Choice {
	s.StatusOrStatementRequested = new(model.StatusOrStatement3Choice)
	return s.StatusOrStatementRequested
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV02) AddProcessingStatus() *model.ProcessingStatus4Choice {
	s.ProcessingStatus = new(model.ProcessingStatus4Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
