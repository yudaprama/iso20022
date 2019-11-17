package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.03 Document"`
	Message *AcceptorReconciliationResponseV03 `xml:"AccptrRcncltnRspn"`
}

func (d *Document01000103) AddMessage() *AcceptorReconciliationResponseV03 {
	d.Message = new(AcceptorReconciliationResponseV03)
	return d.Message
}

// The AcceptorReconciliationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationResponseV03 struct {

	// Reconciliation response message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to thereconciliation response.
	ReconciliationResponse *model.AcceptorReconciliationResponse2 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationResponseV03) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorReconciliationResponseV03) AddReconciliationResponse() *model.AcceptorReconciliationResponse2 {
	a.ReconciliationResponse = new(model.AcceptorReconciliationResponse2)
	return a.ReconciliationResponse
}

func (a *AcceptorReconciliationResponseV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
