package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04100103 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.041.001.03 Document"`
	Message *FundConfirmedCashForecastReportV03 `xml:"FndConfdCshFcstRptV03"`
}

func (d *Document04100103) AddMessage() *FundConfirmedCashForecastReportV03 {
	d.Message = new(FundConfirmedCashForecastReportV03)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundConfirmedCashForecastReport message to the report user, such as an investment manager, to report the confirmed cash incomings and outgoings of one or more investment funds on one or more trade dates.
// The cash movements may result from, for example, redemption, subscription, switch transactions or reinvestment of dividends.
// Usage
// The FundConfirmedCashForecastReport is used to report definitive cash movements, that is it is sent after the cut-off time and/or the price valuation of the fund.
// This message contains incoming and outgoing cash flows that are confirmed, i.e., the price has been applied. If the price is not yet definitive, then the FundEstimatedCashForecastReport message must be used.
// This message allows the report provider to report cash movements in or out of a fund, but does not allow the Sender to categorise these movements, for example by country, or to give details of the underlying orders, commission or charges.
// If the report provider wishes to give detailed information related to cash movements, then the FundDetailedConfirmedCashForecastReport message must be used.
type FundConfirmedCashForecastReportV03 struct {

	// Identifies the message.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example, subscriptions, redemptions or switches to/from a specified investment fund.
	FundCashForecastDetails []*model.FundCashForecast3 `xml:"FndCshFcstDtls"`

	// Estimated net cash as a result of the cash-in and cash-out flows.
	ConsolidatedNetCashForecast *model.NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundConfirmedCashForecastReportV03) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundConfirmedCashForecastReportV03) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundConfirmedCashForecastReportV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV03) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV03) AddMessagePagination() *model.Pagination {
	f.MessagePagination = new(model.Pagination)
	return f.MessagePagination
}

func (f *FundConfirmedCashForecastReportV03) AddFundCashForecastDetails() *model.FundCashForecast3 {
	newValue := new(model.FundCashForecast3)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV03) AddConsolidatedNetCashForecast() *model.NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(model.NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundConfirmedCashForecastReportV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
