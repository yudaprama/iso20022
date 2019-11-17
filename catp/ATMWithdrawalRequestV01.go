package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.001.001.01 Document"`
	Message *ATMWithdrawalRequestV01 `xml:"ATMWdrwlReq"`
}

func (d *Document00100101) AddMessage() *ATMWithdrawalRequestV01 {
	d.Message = new(ATMWithdrawalRequestV01)
	return d.Message
}

// The ATMWithdrawalRequest message is sent by an ATM to an acquirer or its agent to request the approval of a withdrawal transaction at an ATM.
type ATMWithdrawalRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalRequest *model.ContentInformationType10 `xml:"PrtctdATMWdrwlReq,omitempty"`

	// Information related to the request of a withdrawal transaction from an ATM.
	ATMWithdrawalRequest *model.ATMWithdrawalRequest1 `xml:"ATMWdrwlReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalRequestV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMWithdrawalRequestV01) AddProtectedATMWithdrawalRequest() *model.ContentInformationType10 {
	a.ProtectedATMWithdrawalRequest = new(model.ContentInformationType10)
	return a.ProtectedATMWithdrawalRequest
}

func (a *ATMWithdrawalRequestV01) AddATMWithdrawalRequest() *model.ATMWithdrawalRequest1 {
	a.ATMWithdrawalRequest = new(model.ATMWithdrawalRequest1)
	return a.ATMWithdrawalRequest
}

func (a *ATMWithdrawalRequestV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
