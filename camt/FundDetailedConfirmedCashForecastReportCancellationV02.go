package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04500102 struct {
	XMLName xml.Name                                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.045.001.02 Document"`
	Message *FundDetailedConfirmedCashForecastReportCancellationV02 `xml:"FndDtldConfdCshFcstRptCxlV02"`
}

func (d *Document04500102) AddMessage() *FundDetailedConfirmedCashForecastReportCancellationV02 {
	d.Message = new(FundDetailedConfirmedCashForecastReportCancellationV02)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundDetailedConfirmedCashForecastReportCancellation messages to the report user, such as an investment manager, fund accountant or any other interested party, to cancel a previously sent FundDetailedConfirmedCashForecastReport.
// Usage
// The FundDetailedConfirmedCashForecastReportCancellation message is used to cancel an entire FundDetailedConfirmedCashForecastReport message that was previously sent. This message must contain the reference of the message to be cancelled.
// This message may also contain details of the message to be cancelled, but this is not recommended.
type FundDetailedConfirmedCashForecastReportCancellationV02 struct {

	// Identifies the message.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// The FundDetailedConfirmedCashForecastReport to be cancelled.
	CashForecastReportToBeCancelled *model.FundDetailedConfirmedCashForecastReport2 `xml:"CshFcstRptToBeCanc,omitempty"`
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV02) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV02) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV02) AddPreviousReference() *model.AdditionalReference3 {
	f.PreviousReference = new(model.AdditionalReference3)
	return f.PreviousReference
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV02) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV02) AddMessagePagination() *model.Pagination {
	f.MessagePagination = new(model.Pagination)
	return f.MessagePagination
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV02) AddCashForecastReportToBeCancelled() *model.FundDetailedConfirmedCashForecastReport2 {
	f.CashForecastReportToBeCancelled = new(model.FundDetailedConfirmedCashForecastReport2)
	return f.CashForecastReportToBeCancelled
}
