package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Document"`
	Message *ATMExceptionAdviceV01 `xml:"ATMXcptnAdvc"`
}

func (d *Document01100101) AddMessage() *ATMExceptionAdviceV01 {
	d.Message = new(ATMExceptionAdviceV01)
	return d.Message
}

// The ATMExceptionAdvice message is sent by an ATM to an acquirer or its agent to inform of that an exception occurred outside a service.
type ATMExceptionAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMExceptionAdvice *model.ContentInformationType10 `xml:"PrtctdATMXcptnAdvc,omitempty"`

	// Information related to exceptions occurring on the ATM.
	ATMExceptionAdvice *model.ATMExceptionAdvice1 `xml:"ATMXcptnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMExceptionAdviceV01) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMExceptionAdviceV01) AddProtectedATMExceptionAdvice() *model.ContentInformationType10 {
	a.ProtectedATMExceptionAdvice = new(model.ContentInformationType10)
	return a.ProtectedATMExceptionAdvice
}

func (a *ATMExceptionAdviceV01) AddATMExceptionAdvice() *model.ATMExceptionAdvice1 {
	a.ATMExceptionAdvice = new(model.ATMExceptionAdvice1)
	return a.ATMExceptionAdvice
}

func (a *ATMExceptionAdviceV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
