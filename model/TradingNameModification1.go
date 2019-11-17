package model

// Specifies the type of change to the trading name.
type TradingNameModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Name used by a business for commercial purposes, although its registered legal name, used for contracts and other formal situations, may be another.
	TradingName *Max350Text `xml:"TradgNm"`
}

func (t *TradingNameModification1) SetModificationCode(value string) {
	t.ModificationCode = (*Modification1Code)(&value)
}

func (t *TradingNameModification1) SetTradingName(value string) {
	t.TradingName = (*Max350Text)(&value)
}
