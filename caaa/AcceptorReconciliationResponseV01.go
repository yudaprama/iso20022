package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.01 Document"`
	Message *AcceptorReconciliationResponseV01 `xml:"AccptrRcncltnRspn"`
}

func (d *Document01000101) AddMessage() *AcceptorReconciliationResponseV01 {
	d.Message = new(AcceptorReconciliationResponseV01)
	return d.Message
}

// Scope
// The AcceptorReconciliationResponse message is sent by the acquirer to communicate to the card acceptor the totals of the card payment transaction performed for the reconciliation period. An agent never forwards the message.
// Usage
// The AcceptorReconciliationResponse message is used to compare the totals between a card acceptor and an acquirer for the reconciliation period.
type AcceptorReconciliationResponseV01 struct {

	// Reconciliation response message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to thereconciliation response.
	ReconciliationResponse *model.AcceptorReconciliationResponse1 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationResponseV01) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorReconciliationResponseV01) AddReconciliationResponse() *model.AcceptorReconciliationResponse1 {
	a.ReconciliationResponse = new(model.AcceptorReconciliationResponse1)
	return a.ReconciliationResponse
}

func (a *AcceptorReconciliationResponseV01) AddSecurityTrailer() *model.ContentInformationType3 {
	a.SecurityTrailer = new(model.ContentInformationType3)
	return a.SecurityTrailer
}
