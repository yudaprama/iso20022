package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200101 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:cain.012.001.01 Document"`
	Message *KeyExchangeResponse `xml:"KeyXchgRspn"`
}

func (d *Document01200101) AddMessage() *KeyExchangeResponse {
	d.Message = new(KeyExchangeResponse)
	return d.Message
}

// The KeyExchangeResponse message is sent by an acquirer, an issuer or an agent to answer to a KeyExchangeInitiation message and complete a cryptographic key exchange.
type KeyExchangeResponse struct {

	// Information related to the protocol management.
	Header *model.Header17 `xml:"Hdr"`

	// Information related to the response to a key exchange.
	KeyExchangeResponse *model.AcquirerKeyExchangeResponse1 `xml:"KeyXchgRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr"`
}

func (k *KeyExchangeResponse) AddHeader() *model.Header17 {
	k.Header = new(model.Header17)
	return k.Header
}

func (k *KeyExchangeResponse) AddKeyExchangeResponse() *model.AcquirerKeyExchangeResponse1 {
	k.KeyExchangeResponse = new(model.AcquirerKeyExchangeResponse1)
	return k.KeyExchangeResponse
}

func (k *KeyExchangeResponse) AddSecurityTrailer() *model.ContentInformationType12 {
	k.SecurityTrailer = new(model.ContentInformationType12)
	return k.SecurityTrailer
}
