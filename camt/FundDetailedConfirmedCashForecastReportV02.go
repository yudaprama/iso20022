package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04300102 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.043.001.02 Document"`
	Message *FundDetailedConfirmedCashForecastReportV02 `xml:"camt.043.001.02"`
}

func (d *Document04300102) AddMessage() *FundDetailedConfirmedCashForecastReportV02 {
	d.Message = new(FundDetailedConfirmedCashForecastReportV02)
	return d.Message
}

// Scope
// The FundDetailedConfirmedCashForecastReport message is sent by a report provider, such as a transfer agent or a designated agent of the fund, to a report user, such as an investment manager, a fund accountant or any other interested party.
// This message is used to report confirmed cash incomings and outgoings, sorted by country, institution or some other criteria defined by the user. This message can be used to report the confirmed cash movements of one or more investment funds, on one or more trade dates. These cash movements may result from, for example, redemption, subscription, switch transactions or dividends.
// Usage
// The FundDetailedConfirmedCashForecastReport is used to provide definitive cash movements, i.e., it is sent after the cut-off time and/or the price valuation of the fund.
type FundDetailedConfirmedCashForecastReportV02 struct {

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example, subscriptions, redemptions or switches to/from a specified investment fund. The information provided is sorted by pre-defined criteria such as country, institution, currency or user defined criteria.
	FundCashForecastDetails []*model.FundCashForecast2 `xml:"FndCshFcstDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundDetailedConfirmedCashForecastReportV02) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundDetailedConfirmedCashForecastReportV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportV02) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportV02) AddFundCashForecastDetails() *model.FundCashForecast2 {
	newValue := new(model.FundCashForecast2)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
