package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04500101 struct {
	XMLName xml.Name                                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.045.001.01 Document"`
	Message *FundDetailedConfirmedCashForecastReportCancellationV01 `xml:"camt.045.001.01"`
}

func (d *Document04500101) AddMessage() *FundDetailedConfirmedCashForecastReportCancellationV01 {
	d.Message = new(FundDetailedConfirmedCashForecastReportCancellationV01)
	return d.Message
}

// Scope
// The FundDetailedConfirmedCashForecastReportCancellation message is sent by a report provider, such as a transfer agent or a designated agent of the fund, to a report user, such as an investment manager, a fund accountant or any other interested party.
// This message is used to cancel a previously sent FundDetailedConfirmedCashForecastReport message.
// Usage
// The FundDetailedConfirmedCashForecastReportCancellation message is used to cancel an entire FundDetailedConfirmedCashForecastReport message that was previously sent by the report provider.
// This message must contain the reference of the message to be cancelled. This message may also contain details of the message to be cancelled, but this is not recommended.
type FundDetailedConfirmedCashForecastReportCancellationV01 struct {

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// The FundDetailedConfirmedCashForecastReport to be cancelled.
	CashForecastReportToBeCancelled *model.FundDetailedConfirmedCashForecastReport1 `xml:"CshFcstRptToBeCanc,omitempty"`
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV01) AddPoolReference() *model.AdditionalReference3 {
	f.PoolReference = new(model.AdditionalReference3)
	return f.PoolReference
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV01) AddPreviousReference() *model.AdditionalReference3 {
	f.PreviousReference = new(model.AdditionalReference3)
	return f.PreviousReference
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV01) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV01) AddCashForecastReportToBeCancelled() *model.FundDetailedConfirmedCashForecastReport1 {
	f.CashForecastReportToBeCancelled = new(model.FundDetailedConfirmedCashForecastReport1)
	return f.CashForecastReportToBeCancelled
}
