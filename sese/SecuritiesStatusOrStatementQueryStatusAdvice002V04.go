package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200204 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.04 Document"`
	Message *SecuritiesStatusOrStatementQueryStatusAdvice002V04 `xml:"SctiesStsOrStmtQryStsAdvc"`
}

func (d *Document02200204) AddMessage() *SecuritiesStatusOrStatementQueryStatusAdvice002V04 {
	d.Message = new(SecuritiesStatusOrStatementQueryStatusAdvice002V04)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesStatusOrStatementQueryStatusAdvice to an account owner to advise the status of a status query or statement query previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesStatusOrStatementQueryStatusAdvice002V04 struct {

	// Unambiguous identification of the query as per the account owner.
	QueryDetails *model.DocumentIdentification48 `xml:"QryDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	StatusOrStatementRequested *model.StatusOrStatement8Choice `xml:"StsOrStmtReqd,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *model.ProcessingStatus64Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddQueryDetails() *model.DocumentIdentification48 {
	s.QueryDetails = new(model.DocumentIdentification48)
	return s.QueryDetails
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddAccountOwner() *model.PartyIdentification109 {
	s.AccountOwner = new(model.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddSafekeepingAccount() *model.SecuritiesAccount30 {
	s.SafekeepingAccount = new(model.SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddStatusOrStatementRequested() *model.StatusOrStatement8Choice {
	s.StatusOrStatementRequested = new(model.StatusOrStatement8Choice)
	return s.StatusOrStatementRequested
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddProcessingStatus() *model.ProcessingStatus64Choice {
	s.ProcessingStatus = new(model.ProcessingStatus64Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
