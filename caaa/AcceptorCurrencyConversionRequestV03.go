package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600103 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.016.001.03 Document"`
	Message *AcceptorCurrencyConversionRequestV03 `xml:"AccptrCcyConvsReq"`
}

func (d *Document01600103) AddMessage() *AcceptorCurrencyConversionRequestV03 {
	d.Message = new(AcceptorCurrencyConversionRequestV03)
	return d.Message
}

// The AcceptorCurrencyConversionRequest message is sent by the card acceptor to the currency conversion service provider to request if the cardholder is able to pay in the currency of its card.
//
type AcceptorCurrencyConversionRequestV03 struct {

	// Currency Conversion request message management information.
	Header *model.Header30 `xml:"Hdr"`

	// Information related to the currency conversion request.
	CurrencyConversionRequest *model.AcceptorCurrencyConversionRequest3 `xml:"CcyConvsReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCurrencyConversionRequestV03) AddHeader() *model.Header30 {
	a.Header = new(model.Header30)
	return a.Header
}

func (a *AcceptorCurrencyConversionRequestV03) AddCurrencyConversionRequest() *model.AcceptorCurrencyConversionRequest3 {
	a.CurrencyConversionRequest = new(model.AcceptorCurrencyConversionRequest3)
	return a.CurrencyConversionRequest
}

func (a *AcceptorCurrencyConversionRequestV03) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
