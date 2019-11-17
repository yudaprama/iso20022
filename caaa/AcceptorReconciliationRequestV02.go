package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.02 Document"`
	Message *AcceptorReconciliationRequestV02 `xml:"AccptrRcncltnReq"`
}

func (d *Document00900102) AddMessage() *AcceptorReconciliationRequestV02 {
	d.Message = new(AcceptorReconciliationRequestV02)
	return d.Message
}

// The AcceptorReconciliationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationRequestV02 struct {

	// Reconciliation request message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the reconcilliation request.
	ReconciliationRequest *model.AcceptorReconciliationRequest2 `xml:"RcncltnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationRequestV02) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorReconciliationRequestV02) AddReconciliationRequest() *model.AcceptorReconciliationRequest2 {
	a.ReconciliationRequest = new(model.AcceptorReconciliationRequest2)
	return a.ReconciliationRequest
}

func (a *AcceptorReconciliationRequestV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
