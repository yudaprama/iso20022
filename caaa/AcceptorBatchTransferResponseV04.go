package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200104 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.04 Document"`
	Message *AcceptorBatchTransferResponseV04 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200104) AddMessage() *AcceptorBatchTransferResponseV04 {
	d.Message = new(AcceptorBatchTransferResponseV04)
	return d.Message
}

// The AcceptorBatchTransferResponse is sent by the acquirer (or its agent) to inform the acceptor (or its agent) of the transfer in a previous AcceptorBatchTransfer of a collection of transactions.
type AcceptorBatchTransferResponseV04 struct {

	// Capture advice response message management information.
	Header *model.Header12 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	BatchTransferResponse *model.CardPaymentBatchTransferResponse3 `xml:"BtchTrfRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferResponseV04) AddHeader() *model.Header12 {
	a.Header = new(model.Header12)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV04) AddBatchTransferResponse() *model.CardPaymentBatchTransferResponse3 {
	a.BatchTransferResponse = new(model.CardPaymentBatchTransferResponse3)
	return a.BatchTransferResponse
}

func (a *AcceptorBatchTransferResponseV04) AddSecurityTrailer() *model.ContentInformationType12 {
	a.SecurityTrailer = new(model.ContentInformationType12)
	return a.SecurityTrailer
}
