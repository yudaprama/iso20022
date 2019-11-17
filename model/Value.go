package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type Value struct {

	// Specifies the amount in the base currency of the receiver.
	BaseCurrencyItem *ActiveOrHistoricCurrencyAndAmount `xml:"BaseCcyItm"`

	// Specifies the amount in another currency.
	AlternateCurrencyItem []*ActiveOrHistoricCurrencyAndAmount `xml:"AltrnCcyItm"`
}

func (v *Value) SetBaseCurrencyItem(value, currency string) {
	v.BaseCurrencyItem = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (v *Value) AddAlternateCurrencyItem(value, currency string) {
	v.AlternateCurrencyItem = append(v.AlternateCurrencyItem, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}
