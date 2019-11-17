package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200203 struct {
	XMLName xml.Name                                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.002.03 Document"`
	Message *SecuritiesSettlementTransactionAuditTrailReport002V03 `xml:"SctiesSttlmTxAudtTrlRpt"`
}

func (d *Document02200203) AddMessage() *SecuritiesSettlementTransactionAuditTrailReport002V03 {
	d.Message = new(SecuritiesSettlementTransactionAuditTrailReport002V03)
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
type SecuritiesSettlementTransactionAuditTrailReport002V03 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Identification of the SecuritiesStatusQuery message sent to request this report.
	QueryReference *model.Identification16 `xml:"QryRef,omitempty"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications34 `xml:"TxId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	//  Provides the history of status and reasons for a pending, posted or cancelled transaction.
	StatusTrail []*model.StatusTrail7 `xml:"StsTrl,omitempty"`
}

func (s *SecuritiesSettlementTransactionAuditTrailReport002V03) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAuditTrailReport002V03) AddQueryReference() *model.Identification16 {
	s.QueryReference = new(model.Identification16)
	return s.QueryReference
}

func (s *SecuritiesSettlementTransactionAuditTrailReport002V03) AddTransactionIdentification() *model.TransactionIdentifications34 {
	s.TransactionIdentification = new(model.TransactionIdentifications34)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionAuditTrailReport002V03) AddSafekeepingAccount() *model.SecuritiesAccount27 {
	s.SafekeepingAccount = new(model.SecuritiesAccount27)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAuditTrailReport002V03) AddAccountOwner() *model.PartyIdentification109 {
	s.AccountOwner = new(model.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAuditTrailReport002V03) AddStatusTrail() *model.StatusTrail7 {
	newValue := new(model.StatusTrail7)
	s.StatusTrail = append(s.StatusTrail, newValue)
	return newValue
}
