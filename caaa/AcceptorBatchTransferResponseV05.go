package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200105 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.05 Document"`
	Message *AcceptorBatchTransferResponseV05 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200105) AddMessage() *AcceptorBatchTransferResponseV05 {
	d.Message = new(AcceptorBatchTransferResponseV05)
	return d.Message
}

// The AcceptorBatchTransferResponse is sent by the acquirer (or its agent) to inform the acceptor (or its agent) of the transfer in a previous AcceptorBatchTransfer of a collection of transactions.
type AcceptorBatchTransferResponseV05 struct {

	// Capture advice response message management information.
	Header *model.Header25 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	BatchTransferResponse *model.CardPaymentBatchTransferResponse4 `xml:"BtchTrfRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorBatchTransferResponseV05) AddHeader() *model.Header25 {
	a.Header = new(model.Header25)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV05) AddBatchTransferResponse() *model.CardPaymentBatchTransferResponse4 {
	a.BatchTransferResponse = new(model.CardPaymentBatchTransferResponse4)
	return a.BatchTransferResponse
}

func (a *AcceptorBatchTransferResponseV05) AddSecurityTrailer() *model.ContentInformationType12 {
	a.SecurityTrailer = new(model.ContentInformationType12)
	return a.SecurityTrailer
}
