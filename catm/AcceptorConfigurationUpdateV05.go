package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300105 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.05 Document"`
	Message *AcceptorConfigurationUpdateV05 `xml:"AccptrCfgtnUpd"`
}

func (d *Document00300105) AddMessage() *AcceptorConfigurationUpdateV05 {
	d.Message = new(AcceptorConfigurationUpdateV05)
	return d.Message
}

// Update of the acceptor configuration to be downloaded by the terminal management system.
type AcceptorConfigurationUpdateV05 struct {

	// Set of characteristics related to the transfer of the acceptor parameters.
	Header *model.Header27 `xml:"Hdr"`

	// Acceptor configuration to be downloaded from the terminal management system.
	AcceptorConfiguration *model.AcceptorConfiguration5 `xml:"AccptrCfgtn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorConfigurationUpdateV05) AddHeader() *model.Header27 {
	a.Header = new(model.Header27)
	return a.Header
}

func (a *AcceptorConfigurationUpdateV05) AddAcceptorConfiguration() *model.AcceptorConfiguration5 {
	a.AcceptorConfiguration = new(model.AcceptorConfiguration5)
	return a.AcceptorConfiguration
}

func (a *AcceptorConfigurationUpdateV05) AddSecurityTrailer() *model.ContentInformationType12 {
	a.SecurityTrailer = new(model.ContentInformationType12)
	return a.SecurityTrailer
}
