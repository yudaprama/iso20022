package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.003.001.01 Document"`
	Message *AcquirerFinancialInitiation `xml:"AcqrrFinInitn"`
}

func (d *Document00300101) AddMessage() *AcquirerFinancialInitiation {
	d.Message = new(AcquirerFinancialInitiation)
	return d.Message
}

// The AcquirerFinancialInitiation message is sent by an acquirer or  an agent to an issuer or an agent, to request, advice or notify the approval and the clearing of a card transaction.
type AcquirerFinancialInitiation struct {

	// Information related to the protocol management.
	Header *model.Header17 `xml:"Hdr"`

	// Information related to financial authorisation.
	FinancialInitiation *model.AcquirerFinancialInitiation1 `xml:"FinInitn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr"`
}

func (a *AcquirerFinancialInitiation) AddHeader() *model.Header17 {
	a.Header = new(model.Header17)
	return a.Header
}

func (a *AcquirerFinancialInitiation) AddFinancialInitiation() *model.AcquirerFinancialInitiation1 {
	a.FinancialInitiation = new(model.AcquirerFinancialInitiation1)
	return a.FinancialInitiation
}

func (a *AcquirerFinancialInitiation) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
