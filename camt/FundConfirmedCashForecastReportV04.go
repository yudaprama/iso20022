package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04100104 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.041.001.04 Document"`
	Message *FundConfirmedCashForecastReportV04 `xml:"FndConfdCshFcstRpt"`
}

func (d *Document04100104) AddMessage() *FundConfirmedCashForecastReportV04 {
	d.Message = new(FundConfirmedCashForecastReportV04)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundConfirmedCashForecastReport message to the report user, such as an investment manager or pricing agent, to report the confirmed cash incomings and outgoings of one or more share classes of an investment fund on one or more trade dates.
// The cash movements may result from, for example, redemption, subscription, switch transactions or reinvestment of dividends.
// Usage
// The FundConfirmedCashForecastReport is used to report definitive cash movements, that is it is sent after the cut-off time and/or the price valuation of the fund.
// This message contains incoming and outgoing cash flows that are confirmed, that is, the price has been applied. If the price is not yet definitive, then the FundEstimatedCashForecastReport message must be used.
// The message structure allows for the following uses:
// -	to provide cash in and cash out amounts for a fund/sub fund (a FundOrSubFundDetails sequence is used),
// -	to provide cash in and cash out amounts for a fund/sub fund and one or more share classes (FundOrSubFundDetails sequence and one or more FundCashForecastDetails sequences are used),
// -	to provide cash in and cash out amounts for one or more share classes (one or more FundCashForecastDetails sequences are used).
// -	to provide cash in and cash out amounts for more than one fund/sub fund, and more than one share classes (two or more FundOrSubFundDetails sequences and two or more FundCashForecastDetails sequences and used); however, it should be noted that, in this usage, there is no way to determine which share class belongs to which fund/sub fund from the message content itself, which may not be desirable and the use of this kind of combination should be bilaterally agreed.
// This message allows the report provider to report cash movements in or out of a fund, but does not allow the Sender to categorise these movements, for example by country, or to give details of the underlying orders, commission or charges.
// If the report provider wishes to give detailed information related to cash movements, then the FundDetailedConfirmedCashForecastReport message must be used.
type FundConfirmedCashForecastReportV04 struct {

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
	FundOrSubFundDetails []*model.Fund2 `xml:"FndOrSubFndDtls,omitempty"`

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of transactions in shares in an investment fund, for example, subscriptions, redemptions or switches.
	FundCashForecastDetails []*model.FundCashForecast7 `xml:"FndCshFcstDtls,omitempty"`

	// Estimated net cash as a result of the cash-in and cash-out flows.
	ConsolidatedNetCashForecast *model.NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundConfirmedCashForecastReportV04) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundConfirmedCashForecastReportV04) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundConfirmedCashForecastReportV04) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV04) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV04) AddMessagePagination() *model.Pagination {
	f.MessagePagination = new(model.Pagination)
	return f.MessagePagination
}

func (f *FundConfirmedCashForecastReportV04) AddFundOrSubFundDetails() *model.Fund2 {
	newValue := new(model.Fund2)
	f.FundOrSubFundDetails = append(f.FundOrSubFundDetails, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV04) AddFundCashForecastDetails() *model.FundCashForecast7 {
	newValue := new(model.FundCashForecast7)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV04) AddConsolidatedNetCashForecast() *model.NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(model.NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundConfirmedCashForecastReportV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
