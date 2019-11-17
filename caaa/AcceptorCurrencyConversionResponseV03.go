package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700103 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.03 Document"`
	Message *AcceptorCurrencyConversionResponseV03 `xml:"AccptrCcyConvsRspn"`
}

func (d *Document01700103) AddMessage() *AcceptorCurrencyConversionResponseV03 {
	d.Message = new(AcceptorCurrencyConversionResponseV03)
	return d.Message
}

// The AcceptorCurrencyConversionResponse message is sent by currency conversion service provider to the card acceptor to return the result of a potential currency conversion for the cardholder.
//
type AcceptorCurrencyConversionResponseV03 struct {

	// Currency conversion response message management information.
	Header *model.Header30 `xml:"Hdr"`

	// Information related to the outcome of the currency conversion.
	CurrencyConversionResponse *model.AcceptorCurrencyConversionResponse3 `xml:"CcyConvsRspn"`

	// Trailer of the message containing a MAC (message authentication code).
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCurrencyConversionResponseV03) AddHeader() *model.Header30 {
	a.Header = new(model.Header30)
	return a.Header
}

func (a *AcceptorCurrencyConversionResponseV03) AddCurrencyConversionResponse() *model.AcceptorCurrencyConversionResponse3 {
	a.CurrencyConversionResponse = new(model.AcceptorCurrencyConversionResponse3)
	return a.CurrencyConversionResponse
}

func (a *AcceptorCurrencyConversionResponseV03) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
