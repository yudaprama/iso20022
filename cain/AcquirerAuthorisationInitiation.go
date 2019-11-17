package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:cain.001.001.01 Document"`
	Message *AcquirerAuthorisationInitiation `xml:"AcqrrAuthstnInitn"`
}

func (d *Document00100101) AddMessage() *AcquirerAuthorisationInitiation {
	d.Message = new(AcquirerAuthorisationInitiation)
	return d.Message
}

// The AcquirerAuthorisationInitiation message is sent by an acquirer or an agent to an issuer or an agent, to request, advice or notify the approval of a card transaction.
type AcquirerAuthorisationInitiation struct {

	// Information related to the protocol management.
	Header *model.Header17 `xml:"Hdr"`

	// Information related to the authorisation initiation.
	AuthorisationInitiation *model.AcquirerAuthorisationInitiation1 `xml:"AuthstnInitn"`

	// Trailer of the message containing a MAC.
	// It corresponds patially to ISO 8583 field number 53, completed by the field number 64 or 128.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcquirerAuthorisationInitiation) AddHeader() *model.Header17 {
	a.Header = new(model.Header17)
	return a.Header
}

func (a *AcquirerAuthorisationInitiation) AddAuthorisationInitiation() *model.AcquirerAuthorisationInitiation1 {
	a.AuthorisationInitiation = new(model.AcquirerAuthorisationInitiation1)
	return a.AuthorisationInitiation
}

func (a *AcquirerAuthorisationInitiation) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
