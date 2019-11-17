package model

// Specifies target dates dates related to account opening and closing.
type AccountContract4 struct {

	// Date on which the account and related services are expected to cease to be operational for the account owner.
	TargetClosingDate *ISODate `xml:"TrgtClsgDt,omitempty"`

	// Indicator that the account opening/maintenance/closing process needs to be treated urgently, that is, sooner than the terms established by the service level agreed between the account holder customer and the account servicing institution.
	UrgencyFlag *YesNoIndicator `xml:"UrgcyFlg,omitempty"`

	// Indicates removal of the account. After removal, an account will not appear anymore in reports.
	RemovalIndicator *YesNoIndicator `xml:"RmvlInd,omitempty"`
}

func (a *AccountContract4) SetTargetClosingDate(value string) {
	a.TargetClosingDate = (*ISODate)(&value)
}

func (a *AccountContract4) SetUrgencyFlag(value string) {
	a.UrgencyFlag = (*YesNoIndicator)(&value)
}

func (a *AccountContract4) SetRemovalIndicator(value string) {
	a.RemovalIndicator = (*YesNoIndicator)(&value)
}
