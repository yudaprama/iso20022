package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.03 Document"`
	Message *AcceptorReconciliationRequestV03 `xml:"AccptrRcncltnReq"`
}

func (d *Document00900103) AddMessage() *AcceptorReconciliationRequestV03 {
	d.Message = new(AcceptorReconciliationRequestV03)
	return d.Message
}

// The AcceptorReconciliationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationRequestV03 struct {

	// Reconciliation request message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to the reconcilliation request.
	ReconciliationRequest *model.AcceptorReconciliationRequest3 `xml:"RcncltnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationRequestV03) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorReconciliationRequestV03) AddReconciliationRequest() *model.AcceptorReconciliationRequest3 {
	a.ReconciliationRequest = new(model.AcceptorReconciliationRequest3)
	return a.ReconciliationRequest
}

func (a *AcceptorReconciliationRequestV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
