package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100104 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.04 Document"`
	Message *AcceptorBatchTransferV04 `xml:"AccptrBtchTrf"`
}

func (d *Document01100104) AddMessage() *AcceptorBatchTransferV04 {
	d.Message = new(AcceptorBatchTransferV04)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV04 struct {

	// Batch capture message management information.
	Header *model.Header12 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *model.CardPaymentBatchTransfer3 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV04) AddHeader() *model.Header12 {
	a.Header = new(model.Header12)
	return a.Header
}

func (a *AcceptorBatchTransferV04) AddBatchTransfer() *model.CardPaymentBatchTransfer3 {
	a.BatchTransfer = new(model.CardPaymentBatchTransfer3)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV04) AddSecurityTrailer() *model.ContentInformationType12 {
	a.SecurityTrailer = new(model.ContentInformationType12)
	return a.SecurityTrailer
}
