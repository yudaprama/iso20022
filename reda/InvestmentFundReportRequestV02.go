package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.02 Document"`
	Message *InvestmentFundReportRequestV02 `xml:"InvstmtFndRptReq"`
}

func (d *Document00500102) AddMessage() *InvestmentFundReportRequestV02 {
	d.Message = new(InvestmentFundReportRequestV02)
	return d.Message
}

// Scope
// A report user, for example, a professional investor, investment fund distributor, market data provider, regulator or other interested party sends the InvestmentFundReportRequest message to the report provider, for example, a fund promoter, fund management company, transfer agent, or market data provider to request a report.
// The InvestmentFundReportRequest message can be used to request one or many fund processing passport reports.
// Usage
// If the InvestmentFundReportRequest message is used to request a fund processing passport then the request can specify the financial instrument for which the report is requested. Other appropriate parameters can also be included. It is also possible to indicate that the request is an open request, that is, there is no specific criteria for the report requested. For example, a request for a fund processing passport report that is specified as "no criteria" means that the request is a request for all fund processing passports.
type InvestmentFundReportRequestV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Parameters for which the fund processing passport report is requested.
	FundProcessingPassportReport []*model.FundParameters3Choice `xml:"FPPRpt"`
}

func (i *InvestmentFundReportRequestV02) AddMessageIdentification() *model.MessageIdentification1 {
	i.MessageIdentification = new(model.MessageIdentification1)
	return i.MessageIdentification
}

func (i *InvestmentFundReportRequestV02) AddPreviousReference() *model.AdditionalReference3 {
	i.PreviousReference = new(model.AdditionalReference3)
	return i.PreviousReference
}

func (i *InvestmentFundReportRequestV02) AddRelatedReference() *model.AdditionalReference3 {
	i.RelatedReference = new(model.AdditionalReference3)
	return i.RelatedReference
}

func (i *InvestmentFundReportRequestV02) AddFundProcessingPassportReport() *model.FundParameters3Choice {
	newValue := new(model.FundParameters3Choice)
	i.FundProcessingPassportReport = append(i.FundProcessingPassportReport, newValue)
	return newValue
}
