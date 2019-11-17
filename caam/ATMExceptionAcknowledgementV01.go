package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.012.001.01 Document"`
	Message *ATMExceptionAcknowledgementV01 `xml:"ATMXcptnAck"`
}

func (d *Document01200101) AddMessage() *ATMExceptionAcknowledgementV01 {
	d.Message = new(ATMExceptionAcknowledgementV01)
	return d.Message
}

// The ATMExceptionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMExceptionAdvice message.
type ATMExceptionAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMExceptionAcknowledgement *model.ContentInformationType10 `xml:"PrtctdATMXcptnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM exception.
	ATMExceptionAcknowledgement *model.ATMExceptionAcknowledgement1 `xml:"ATMXcptnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMExceptionAcknowledgementV01) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMExceptionAcknowledgementV01) AddProtectedATMExceptionAcknowledgement() *model.ContentInformationType10 {
	a.ProtectedATMExceptionAcknowledgement = new(model.ContentInformationType10)
	return a.ProtectedATMExceptionAcknowledgement
}

func (a *ATMExceptionAcknowledgementV01) AddATMExceptionAcknowledgement() *model.ATMExceptionAcknowledgement1 {
	a.ATMExceptionAcknowledgement = new(model.ATMExceptionAcknowledgement1)
	return a.ATMExceptionAcknowledgement
}

func (a *ATMExceptionAcknowledgementV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
