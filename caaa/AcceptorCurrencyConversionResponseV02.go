package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700102 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.02 Document"`
	Message *AcceptorCurrencyConversionResponseV02 `xml:"AccptrCcyConvsRspn"`
}

func (d *Document01700102) AddMessage() *AcceptorCurrencyConversionResponseV02 {
	d.Message = new(AcceptorCurrencyConversionResponseV02)
	return d.Message
}

// The AcceptorCurrencyConversionResponse message is sent by currency conversion service provider to the card acceptor to return the result of a potential currency conversion for the cardholder.
//
type AcceptorCurrencyConversionResponseV02 struct {

	// Currency conversion response message management information.
	Header *model.Header10 `xml:"Hdr"`

	// Information related to the outcome of the currency conversion.
	CurrencyConversionResponse *model.AcceptorCurrencyConversionResponse2 `xml:"CcyConvsRspn"`

	// Trailer of the message containing a MAC (message authentication code).
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCurrencyConversionResponseV02) AddHeader() *model.Header10 {
	a.Header = new(model.Header10)
	return a.Header
}

func (a *AcceptorCurrencyConversionResponseV02) AddCurrencyConversionResponse() *model.AcceptorCurrencyConversionResponse2 {
	a.CurrencyConversionResponse = new(model.AcceptorCurrencyConversionResponse2)
	return a.CurrencyConversionResponse
}

func (a *AcceptorCurrencyConversionResponseV02) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
