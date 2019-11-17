package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200103 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Document"`
	Message *SecuritiesSettlementTransactionAuditTrailReportV03 `xml:"SctiesSttlmTxAudtTrlRpt"`
}

func (d *Document02200103) AddMessage() *SecuritiesSettlementTransactionAuditTrailReportV03 {
	d.Message = new(SecuritiesSettlementTransactionAuditTrailReportV03)
	return d.Message
}

// Scope
// This message is sent by the Market Infrastructure to the CSD to advise of the history of all the statuses, modifications, replacement and cancellation of a specific transaction during its whole life cycle when the instructing party is a direct participant to the Settlement Infrastructure.
//
// Usage
// The message may also be used to:
// - re-send a message sent by the market infrastructure to the direct participant,
// - provide a third party with a copy of a message being sent by the market infrastructure for information,
// - re-send to a third party a copy of a message being sent by the market infrastructure for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionAuditTrailReportV03 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Identification of the SecuritiesStatusQuery message sent to request this report.
	QueryReference *model.Identification14 `xml:"QryRef,omitempty"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications29 `xml:"TxId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	//  Provides the history of status and reasons for a pending, posted or cancelled transaction.
	StatusTrail []*model.StatusTrail6 `xml:"StsTrl,omitempty"`
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddQueryReference() *model.Identification14 {
	s.QueryReference = new(model.Identification14)
	return s.QueryReference
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddTransactionIdentification() *model.TransactionIdentifications29 {
	s.TransactionIdentification = new(model.TransactionIdentifications29)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddSafekeepingAccount() *model.SecuritiesAccount24 {
	s.SafekeepingAccount = new(model.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddStatusTrail() *model.StatusTrail6 {
	newValue := new(model.StatusTrail6)
	s.StatusTrail = append(s.StatusTrail, newValue)
	return newValue
}
