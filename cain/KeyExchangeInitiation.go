package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:cain.011.001.01 Document"`
	Message *KeyExchangeInitiation `xml:"KeyXchgInitn"`
}

func (d *Document01100101) AddMessage() *KeyExchangeInitiation {
	d.Message = new(KeyExchangeInitiation)
	return d.Message
}

// The KeyExchangeInitiation message is sent by any party to an acquirer, an issuer or an agent, to initiate a cryptographic key exchange.
type KeyExchangeInitiation struct {

	// Information related to the protocol management.
	Header *model.Header17 `xml:"Hdr"`

	// Information related to the key exchange.
	KeyExchangeInitiation *model.AcquirerKeyExchangeInitiation1 `xml:"KeyXchgInitn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr"`
}

func (k *KeyExchangeInitiation) AddHeader() *model.Header17 {
	k.Header = new(model.Header17)
	return k.Header
}

func (k *KeyExchangeInitiation) AddKeyExchangeInitiation() *model.AcquirerKeyExchangeInitiation1 {
	k.KeyExchangeInitiation = new(model.AcquirerKeyExchangeInitiation1)
	return k.KeyExchangeInitiation
}

func (k *KeyExchangeInitiation) AddSecurityTrailer() *model.ContentInformationType12 {
	k.SecurityTrailer = new(model.ContentInformationType12)
	return k.SecurityTrailer
}
