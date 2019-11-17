package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:cain.005.001.01 Document"`
	Message *AcquirerReversalInitiation `xml:"AcqrrRvslInitn"`
}

func (d *Document00500101) AddMessage() *AcquirerReversalInitiation {
	d.Message = new(AcquirerReversalInitiation)
	return d.Message
}

// The AcquirerReversalInitiation message is sent by an acquirer or an agent to an issuer or an agent, to request, advice or notify the reversal of a card transaction.
type AcquirerReversalInitiation struct {

	// Information related to the protocol management
	Header *model.Header18 `xml:"Hdr"`

	// Information related to the reversal.
	ReversalInitiation *model.AcquirerReversalInitiation1 `xml:"RvslInitn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcquirerReversalInitiation) AddHeader() *model.Header18 {
	a.Header = new(model.Header18)
	return a.Header
}

func (a *AcquirerReversalInitiation) AddReversalInitiation() *model.AcquirerReversalInitiation1 {
	a.ReversalInitiation = new(model.AcquirerReversalInitiation1)
	return a.ReversalInitiation
}

func (a *AcquirerReversalInitiation) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
