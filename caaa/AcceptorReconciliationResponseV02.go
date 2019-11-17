package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000102 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.02 Document"`
	Message *AcceptorReconciliationResponseV02 `xml:"AccptrRcncltnRspn"`
}

func (d *Document01000102) AddMessage() *AcceptorReconciliationResponseV02 {
	d.Message = new(AcceptorReconciliationResponseV02)
	return d.Message
}

// The AcceptorReconciliationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationResponseV02 struct {

	// Reconciliation response message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to thereconciliation response.
	ReconciliationResponse *model.AcceptorReconciliationResponse2 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationResponseV02) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorReconciliationResponseV02) AddReconciliationResponse() *model.AcceptorReconciliationResponse2 {
	a.ReconciliationResponse = new(model.AcceptorReconciliationResponse2)
	return a.ReconciliationResponse
}

func (a *AcceptorReconciliationResponseV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
