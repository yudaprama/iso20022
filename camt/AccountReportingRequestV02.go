package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document06000102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.02 Document"`
	Message *AccountReportingRequestV02 `xml:"AcctRptgReq"`
}

func (d *Document06000102) AddMessage() *AccountReportingRequestV02 {
	d.Message = new(AccountReportingRequestV02)
	return d.Message
}

// Scope
// The AccountReportingRequest message is sent by the account owner, either directly or through a forwarding agent, to one of its account servicing institutions. It is used to ask the account servicing institution to send a report on the account owner's account in a BankToCustomerAccountReport (camt.052.001.02), a BankToCustomerStatement (camt.053.001.02) or a BankToCustomerDebitCreditNotification (camt.054.001.02).
// Usage
// The AccountReportingRequest message is used to advise the account servicing institution of funds that the account owner expects to have credited to its account. The message can be used in either a direct or a relay scenario.
type AccountReportingRequestV02 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *model.GroupHeader43 `xml:"GrpHdr"`

	// Set of elements used to provide further details on the reporting request.
	ReportingRequest []*model.ReportingRequest2 `xml:"RptgReq"`
}

func (a *AccountReportingRequestV02) AddGroupHeader() *model.GroupHeader43 {
	a.GroupHeader = new(model.GroupHeader43)
	return a.GroupHeader
}

func (a *AccountReportingRequestV02) AddReportingRequest() *model.ReportingRequest2 {
	newValue := new(model.ReportingRequest2)
	a.ReportingRequest = append(a.ReportingRequest, newValue)
	return newValue
}
