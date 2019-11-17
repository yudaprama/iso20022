package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.01 Document"`
	Message *AcceptorCurrencyConversionResponseV01 `xml:"AccptrCcyConvsRspn"`
}

func (d *Document01700101) AddMessage() *AcceptorCurrencyConversionResponseV01 {
	d.Message = new(AcceptorCurrencyConversionResponseV01)
	return d.Message
}

// The AcceptorCurrencyConversionResponse message is sent by currency conversion service provider to the card acceptor to return the result of a potential currency conversion for the cardholder.
//
type AcceptorCurrencyConversionResponseV01 struct {

	// Currency conversion response message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to the outcome of the currency conversion.
	CurrencyConversionResponse *model.AcceptorCurrencyConversionResponse1 `xml:"CcyConvsRspn"`

	// Trailer of the message containing a MAC (message authentication code).
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCurrencyConversionResponseV01) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorCurrencyConversionResponseV01) AddCurrencyConversionResponse() *model.AcceptorCurrencyConversionResponse1 {
	a.CurrencyConversionResponse = new(model.AcceptorCurrencyConversionResponse1)
	return a.CurrencyConversionResponse
}

func (a *AcceptorCurrencyConversionResponseV01) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
