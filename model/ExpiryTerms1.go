package model

// Expiry conditions.
type ExpiryTerms1 struct {

	// Date and time when the undertaking will cease to be available.
	DateTime *DateAndDateTimeChoice `xml:"DtTm,omitempty"`

	// Details related to the automatic extension of the undertaking.
	AutoExtension *AutoExtension1 `xml:"AutoXtnsn,omitempty"`

	// Documentary condition that indicates when the undertaking will cease to be available.
	Condition *Max2000Text `xml:"Cond,omitempty"`

	// Indicates whether the expiry terms are without a fixed expiry date.
	OpenEndedIndicator *YesNoIndicator `xml:"OpnEnddInd,omitempty"`
}

func (e *ExpiryTerms1) AddDateTime() *DateAndDateTimeChoice {
	e.DateTime = new(DateAndDateTimeChoice)
	return e.DateTime
}

func (e *ExpiryTerms1) AddAutoExtension() *AutoExtension1 {
	e.AutoExtension = new(AutoExtension1)
	return e.AutoExtension
}

func (e *ExpiryTerms1) SetCondition(value string) {
	e.Condition = (*Max2000Text)(&value)
}

func (e *ExpiryTerms1) SetOpenEndedIndicator(value string) {
	e.OpenEndedIndicator = (*YesNoIndicator)(&value)
}
