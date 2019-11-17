package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900104 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.04 Document"`
	Message *AcceptorReconciliationRequestV04 `xml:"AccptrRcncltnReq"`
}

func (d *Document00900104) AddMessage() *AcceptorReconciliationRequestV04 {
	d.Message = new(AcceptorReconciliationRequestV04)
	return d.Message
}

// The AcceptorReconciliationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationRequestV04 struct {

	// Reconciliation request message management information.
	Header *model.Header10 `xml:"Hdr"`

	// Information related to the reconciliation request.
	ReconciliationRequest *model.AcceptorReconciliationRequest4 `xml:"RcncltnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationRequestV04) AddHeader() *model.Header10 {
	a.Header = new(model.Header10)
	return a.Header
}

func (a *AcceptorReconciliationRequestV04) AddReconciliationRequest() *model.AcceptorReconciliationRequest4 {
	a.ReconciliationRequest = new(model.AcceptorReconciliationRequest4)
	return a.ReconciliationRequest
}

func (a *AcceptorReconciliationRequestV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
