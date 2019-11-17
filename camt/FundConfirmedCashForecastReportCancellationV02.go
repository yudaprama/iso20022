package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04400102 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.044.001.02 Document"`
	Message *FundConfirmedCashForecastReportCancellationV02 `xml:"FndConfdCshFcstRptCxlV02"`
}

func (d *Document04400102) AddMessage() *FundConfirmedCashForecastReportCancellationV02 {
	d.Message = new(FundConfirmedCashForecastReportCancellationV02)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundConfirmedCashForecastReportCancellation message to the report user, such as an investment manager, to cancel a previously sent FundConfirmedCashForecastReport message.
// Usage
// The FundConfirmedCashForecastReportCancellation message is used to cancel an entire FundConfirmedCashForecastReport message that was previously sent by the report provider. This message must contain reference to the of the message being cancelled.
// This message may also contain details of the message to be cancelled, but this is not recommended.
type FundConfirmedCashForecastReportCancellationV02 struct {

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
	CashForecastReportToBeCancelled *model.FundConfirmedCashForecastReport2 `xml:"CshFcstRptToBeCanc,omitempty"`
}

func (f *FundConfirmedCashForecastReportCancellationV02) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundConfirmedCashForecastReportCancellationV02) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundConfirmedCashForecastReportCancellationV02) AddPreviousReference() *model.AdditionalReference3 {
	f.PreviousReference = new(model.AdditionalReference3)
	return f.PreviousReference
}

func (f *FundConfirmedCashForecastReportCancellationV02) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportCancellationV02) AddMessagePagination() *model.Pagination {
	f.MessagePagination = new(model.Pagination)
	return f.MessagePagination
}

func (f *FundConfirmedCashForecastReportCancellationV02) AddCashForecastReportToBeCancelled() *model.FundConfirmedCashForecastReport2 {
	f.CashForecastReportToBeCancelled = new(model.FundConfirmedCashForecastReport2)
	return f.CashForecastReportToBeCancelled
}
