package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04200104 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.042.001.04 Document"`
	Message *FundDetailedEstimatedCashForecastReportV04 `xml:"FndDtldEstmtdCshFcstRpt"`
}

func (d *Document04200104) AddMessage() *FundDetailedEstimatedCashForecastReportV04 {
	d.Message = new(FundDetailedEstimatedCashForecastReportV04)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundDetailedEstimatedCashForecastReport message to the report user, such as an investment manager or pricing agent, to report the estimated cash incomings and outgoings, sorted by country, institution name or other criteria defined by the user of one or more share classes of an investment fund on one or more trade dates.
// The cash movements may result from, for example, redemption, subscription, switch transactions or reinvestment of dividends.
// Usage
// The FundDetailedEstimatedCashForecastReport is used to provide estimated cash movements, that is, it is sent prior to the cut-off time and/or the price valuation of the fund. The message contains incoming and outgoing cash flows that are estimated, that is, the price has not been applied. If the price is definitive, then the FundDetailedConfirmedCashForecastReport message must be used.
// The message structure allows for the following uses:
// -	to provide cash in and cash out amounts for a fund/sub fund and one or more share classes (a FundOrSubFundDetails sequence and one or more EstimatedFundCashForecastDetails sequences are used),
// -	to provide cash in and cash out amounts for one or more share classes (one or more EstimatedFundCashForecastDetails sequences are used).
// If the report is to provide estimated cash in and cash out for a fund/sub fund only and not for one or more share classes, then the FundEstimatedCashForecastReport message must be used.
// The FundDetailedEstimatedCashForecastReport message is used to report cash movements in or out of a fund, organised by party, such as fund management company, country, currency or by some other criteria defined by the report provider. If the report is used to give the cash-in and cash-out for a party, then additional criteria, such as currency and country, can be specified.
// In addition, the underlying transaction type for the cash-in or cash-out movement can be specified, as well as information about the cash movement's underlying orders, such as commission and charges.
type FundDetailedEstimatedCashForecastReportV04 struct {

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

	// Information about the fund/sub fund when the report either specifies cash flow for the fund/sub fund or for a share class of the fund/sub fund.
	FundOrSubFundDetails *model.Fund3 `xml:"FndOrSubFndDtls,omitempty"`

	// Information related to the estimated cash-in and cash-out flows for a specific trade date as a result of transactions in shares in an investment fund, for example, subscriptions, redemptions or switches. The information provided is sorted by pre-defined criteria such as country, institution, currency or user defined criteria.
	EstimatedFundCashForecastDetails []*model.EstimatedFundCashForecast5 `xml:"EstmtdFndCshFcstDtls"`

	// Estimated net cash as a result of the cash-in and cash-out flows specified in the fund cash forecast details.
	ConsolidatedNetCashForecast *model.NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddMessagePagination() *model.Pagination {
	f.MessagePagination = new(model.Pagination)
	return f.MessagePagination
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddFundOrSubFundDetails() *model.Fund3 {
	f.FundOrSubFundDetails = new(model.Fund3)
	return f.FundOrSubFundDetails
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddEstimatedFundCashForecastDetails() *model.EstimatedFundCashForecast5 {
	newValue := new(model.EstimatedFundCashForecast5)
	f.EstimatedFundCashForecastDetails = append(f.EstimatedFundCashForecastDetails, newValue)
	return newValue
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddConsolidatedNetCashForecast() *model.NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(model.NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
