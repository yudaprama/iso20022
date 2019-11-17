package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.02 Document"`
	Message *AcceptorConfigurationUpdateV02 `xml:"AccptrCfgtnUpd"`
}

func (d *Document00300102) AddMessage() *AcceptorConfigurationUpdateV02 {
	d.Message = new(AcceptorConfigurationUpdateV02)
	return d.Message
}

// Update of the acceptor configuration to be dowloaded by the terminal management system.
type AcceptorConfigurationUpdateV02 struct {

	// Set of characteristics related to the transfer of the acceptor parameters.
	Header *model.Header4 `xml:"Hdr"`

	// Acceptor configuration to be downloaded from the terminal management system.
	AcceptorConfiguration *model.AcceptorConfiguration2 `xml:"AccptrCfgtn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType4 `xml:"SctyTrlr"`
}

func (a *AcceptorConfigurationUpdateV02) AddHeader() *model.Header4 {
	a.Header = new(model.Header4)
	return a.Header
}

func (a *AcceptorConfigurationUpdateV02) AddAcceptorConfiguration() *model.AcceptorConfiguration2 {
	a.AcceptorConfiguration = new(model.AcceptorConfiguration2)
	return a.AcceptorConfiguration
}

func (a *AcceptorConfigurationUpdateV02) AddSecurityTrailer() *model.ContentInformationType4 {
	a.SecurityTrailer = new(model.ContentInformationType4)
	return a.SecurityTrailer
}
