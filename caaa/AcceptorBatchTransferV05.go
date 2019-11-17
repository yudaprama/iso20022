package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100105 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.05 Document"`
	Message *AcceptorBatchTransferV05 `xml:"AccptrBtchTrf"`
}

func (d *Document01100105) AddMessage() *AcceptorBatchTransferV05 {
	d.Message = new(AcceptorBatchTransferV05)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV05 struct {

	// Batch capture message management information.
	Header *model.Header25 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *model.CardPaymentBatchTransfer4 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorBatchTransferV05) AddHeader() *model.Header25 {
	a.Header = new(model.Header25)
	return a.Header
}

func (a *AcceptorBatchTransferV05) AddBatchTransfer() *model.CardPaymentBatchTransfer4 {
	a.BatchTransfer = new(model.CardPaymentBatchTransfer4)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV05) AddSecurityTrailer() *model.ContentInformationType12 {
	a.SecurityTrailer = new(model.ContentInformationType12)
	return a.SecurityTrailer
}
