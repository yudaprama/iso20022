package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200101 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.01 Document"`
	Message *SecuritiesSettlementTransactionAuditTrailReportV01 `xml:"SctiesSttlmTxAudtTrlRpt"`
}

func (d *Document02200101) AddMessage() *SecuritiesSettlementTransactionAuditTrailReportV01 {
	d.Message = new(SecuritiesSettlementTransactionAuditTrailReportV01)
	return d.Message
}

// Scope
//
// This message is sent by the Market Infrastructure to the CSD to advise of the history of all the statuses, modifications, replacement and cancellation of a specific transaction during its whole life cycle when the instructing party is a direct participant to the Settlement Infrastructure.
//
//
// Usage
//
// The message may also be used to:
//
// - re-send a message sent by the market infrastructure to the direct participant,
//
// - provide a third party with a copy of a message being sent by the market infrastructure for information,
//
// - re-send to a third party a copy of a message being sent by the market infrastructure for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionAuditTrailReportV01 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Identification of the SecuritiesStatusQuery message sent to request this report.
	QueryReference *model.Identification1 `xml:"QryRef,omitempty"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications15 `xml:"TxId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	//  Provides the history of status and reasons for a pending, posted or cancelled transaction.
	StatusTrail []*model.StatusTrail2 `xml:"StsTrl,omitempty"`
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV01) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV01) AddQueryReference() *model.Identification1 {
	s.QueryReference = new(model.Identification1)
	return s.QueryReference
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV01) AddTransactionIdentification() *model.TransactionIdentifications15 {
	s.TransactionIdentification = new(model.TransactionIdentifications15)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV01) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV01) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV01) AddStatusTrail() *model.StatusTrail2 {
	newValue := new(model.StatusTrail2)
	s.StatusTrail = append(s.StatusTrail, newValue)
	return newValue
}
