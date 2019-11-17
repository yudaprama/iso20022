package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.02 Document"`
	Message *AcceptorBatchTransferV02 `xml:"AccptrBtchTrf"`
}

func (d *Document01100102) AddMessage() *AcceptorBatchTransferV02 {
	d.Message = new(AcceptorBatchTransferV02)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV02 struct {

	// Batch capture message management information.
	Header *model.Header3 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *model.CardPaymentBatchTransfer1 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType4 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV02) AddHeader() *model.Header3 {
	a.Header = new(model.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferV02) AddBatchTransfer() *model.CardPaymentBatchTransfer1 {
	a.BatchTransfer = new(model.CardPaymentBatchTransfer1)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV02) AddSecurityTrailer() *model.ContentInformationType4 {
	a.SecurityTrailer = new(model.ContentInformationType4)
	return a.SecurityTrailer
}
