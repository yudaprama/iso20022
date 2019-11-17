package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04300103 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.043.001.03 Document"`
	Message *FundDetailedConfirmedCashForecastReportV03 `xml:"FndDtldConfdCshFcstRptV03"`
}

func (d *Document04300103) AddMessage() *FundDetailedConfirmedCashForecastReportV03 {
	d.Message = new(FundDetailedConfirmedCashForecastReportV03)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundDetailedConfirmedCashForecastReport message to the report user, such as an investment manager, to report the confirmed cash incomings and outgoings, sorted by country, institution name or other criteria defined by the user of one or more investment funds on one or more trade dates.
// The cash movements may result from, for example, redemption, subscription, switch transactions or reinvestment of dividends.
// Usage
// The FundDetailedConfirmedCashForecastReport is used to provide definitive cash movements, i.e., it is sent after the cut-off time and/or the price valuation of the fund. If the price is not yet definitive, then the FundDetailedEstimatedCashForecastReport message must be used.
// The FundDetailedConfirmedCashForecastReport message is used to report cash movements in or out of a fund, organised by party, such as fund management company, country, currency or by some other criteria defined by the report provider. If the report is used to given the cash-in and cash-out for a party, then additional criteria, such as currency and country, can be specified.
// In addition, the underlying transaction type for the cash-in or cash-out movement can be specified, as well as information about the cash movement's underlying orders, such as commission and charges.
type FundDetailedConfirmedCashForecastReportV03 struct {

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

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example, subscriptions, redemptions or switches to/from a specified investment fund. The information provided is sorted by pre-defined criteria such as country, institution, currency or user defined criteria.
	FundCashForecastDetails []*model.FundCashForecast4 `xml:"FndCshFcstDtls"`

	// Net cash as a result of the cash-in and cash-out flows specified in the fund cash forecast details.
	ConsolidatedNetCashForecast *model.NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddMessagePagination() *model.Pagination {
	f.MessagePagination = new(model.Pagination)
	return f.MessagePagination
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddFundCashForecastDetails() *model.FundCashForecast4 {
	newValue := new(model.FundCashForecast4)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddConsolidatedNetCashForecast() *model.NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(model.NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundDetailedConfirmedCashForecastReportV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
