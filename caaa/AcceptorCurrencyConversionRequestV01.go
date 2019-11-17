package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600101 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.016.001.01 Document"`
	Message *AcceptorCurrencyConversionRequestV01 `xml:"AccptrCcyConvsReq"`
}

func (d *Document01600101) AddMessage() *AcceptorCurrencyConversionRequestV01 {
	d.Message = new(AcceptorCurrencyConversionRequestV01)
	return d.Message
}

// The AcceptorCurrencyConversionRequest message is sent by the card acceptor to the currency conversion service provider to request if the cardholder is able to pay in the currency of its card.
//
type AcceptorCurrencyConversionRequestV01 struct {

	// Currency Conversion request message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to the currency conversion request.
	CurrencyConversionRequest *model.AcceptorCurrencyConversionRequest1 `xml:"CcyConvsReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCurrencyConversionRequestV01) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorCurrencyConversionRequestV01) AddCurrencyConversionRequest() *model.AcceptorCurrencyConversionRequest1 {
	a.CurrencyConversionRequest = new(model.AcceptorCurrencyConversionRequest1)
	return a.CurrencyConversionRequest
}

func (a *AcceptorCurrencyConversionRequestV01) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
