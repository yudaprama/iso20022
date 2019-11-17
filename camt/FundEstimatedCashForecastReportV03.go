package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04000103 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.040.001.03 Document"`
	Message *FundEstimatedCashForecastReportV03 `xml:"FndEstmtdCshFcstRptV03"`
}

func (d *Document04000103) AddMessage() *FundEstimatedCashForecastReportV03 {
	d.Message = new(FundEstimatedCashForecastReportV03)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundEstimatedCashForecastReport message to the report user, such as an investment manager, to report the estimated cash incomings and outgoings of one or more investment funds on one or more trade dates.
// The cash movements may result from, for example, redemption, subscription, switch transactions or reinvestment of dividends.
// Usage
// The FundEstimatedCashForecastReport is used to report estimated cash movements, i.e., it is sent prior to the cut-off time and/or the price valuation of the fund.
// The FundEstimatedCashForecastReport contains incoming and outgoing cash flows that are estimated, i.e., the price has not been applied. If the price is definitive, then the FundConfirmedCashForecastReport message must be used.
// This message allows the report provider to report estimated cash movements in or out of a fund, but does not allow the Sender to categorise these movements, for example by country, or to give details of the underlying orders, commission or charges. If the report provider wishes to give detailed information related to estimated cash movements, then the FundDetailedEstimatedCashForecastReport message must be used.
type FundEstimatedCashForecastReportV03 struct {

	// Identifies the message.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// Information related to the estimated cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example,  subscriptions, redemptions or switch to/from a specified investment fund.
	EstimatedFundCashForecastDetails []*model.EstimatedFundCashForecast3 `xml:"EstmtdFndCshFcstDtls"`

	// Estimated net cash as a result of the cash-in and cash-out flows.
	ConsolidatedNetCashForecast *model.NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundEstimatedCashForecastReportV03) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundEstimatedCashForecastReportV03) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundEstimatedCashForecastReportV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundEstimatedCashForecastReportV03) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundEstimatedCashForecastReportV03) AddMessagePagination() *model.Pagination {
	f.MessagePagination = new(model.Pagination)
	return f.MessagePagination
}

func (f *FundEstimatedCashForecastReportV03) AddEstimatedFundCashForecastDetails() *model.EstimatedFundCashForecast3 {
	newValue := new(model.EstimatedFundCashForecast3)
	f.EstimatedFundCashForecastDetails = append(f.EstimatedFundCashForecastDetails, newValue)
	return newValue
}

func (f *FundEstimatedCashForecastReportV03) AddConsolidatedNetCashForecast() *model.NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(model.NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundEstimatedCashForecastReportV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
