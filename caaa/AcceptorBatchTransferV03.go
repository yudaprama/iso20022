package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.03 Document"`
	Message *AcceptorBatchTransferV03 `xml:"AccptrBtchTrf"`
}

func (d *Document01100103) AddMessage() *AcceptorBatchTransferV03 {
	d.Message = new(AcceptorBatchTransferV03)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV03 struct {

	// Batch capture message management information.
	Header *model.Header3 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *model.CardPaymentBatchTransfer2 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType9 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV03) AddHeader() *model.Header3 {
	a.Header = new(model.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferV03) AddBatchTransfer() *model.CardPaymentBatchTransfer2 {
	a.BatchTransfer = new(model.CardPaymentBatchTransfer2)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV03) AddSecurityTrailer() *model.ContentInformationType9 {
	a.SecurityTrailer = new(model.ContentInformationType9)
	return a.SecurityTrailer
}
