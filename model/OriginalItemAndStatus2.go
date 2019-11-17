package model

// Identify the original notification item and to provide the status.
type OriginalItemAndStatus2 struct {

	// Identification of the original notification item.
	OriginalItemIdentification *Max35Text `xml:"OrgnlItmId"`

	// Unique identification as assigned by the debtor to unambiguously identify the original underlying transaction to the creditor.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money expected to be credited to the account, as per the original notification to receive.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Value date on which the account was expected to be credited.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Specifies the status of the notification item.
	ItemStatus *NotificationStatus3Code `xml:"ItmSts"`

	// Further details of the item status.
	AdditionalStatusInformation *Max105Text `xml:"AddtlStsInf,omitempty"`

	// Provides further information in order to identify a previous payment notification.
	OriginalItemReference *OriginalItemReference1 `xml:"OrgnlItmRef,omitempty"`
}

func (o *OriginalItemAndStatus2) SetOriginalItemIdentification(value string) {
	o.OriginalItemIdentification = (*Max35Text)(&value)
}

func (o *OriginalItemAndStatus2) SetOriginalEndToEndIdentification(value string) {
	o.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (o *OriginalItemAndStatus2) SetAmount(value, currency string) {
	o.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalItemAndStatus2) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalItemAndStatus2) SetItemStatus(value string) {
	o.ItemStatus = (*NotificationStatus3Code)(&value)
}

func (o *OriginalItemAndStatus2) SetAdditionalStatusInformation(value string) {
	o.AdditionalStatusInformation = (*Max105Text)(&value)
}

func (o *OriginalItemAndStatus2) AddOriginalItemReference() *OriginalItemReference1 {
	o.OriginalItemReference = new(OriginalItemReference1)
	return o.OriginalItemReference
}
