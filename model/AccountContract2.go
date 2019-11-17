package model

// Specifies target dates dates related to account opening and closing.
type AccountContract2 struct {

	// Date on which the account and related basic services  are expected to be operational for the account owner.
	TargetGoLiveDate *ISODate `xml:"TrgtGoLiveDt,omitempty"`

	// Date on which the account and related services are expected to cease to be operational for the account owner.
	TargetClosingDate *ISODate `xml:"TrgtClsgDt,omitempty"`

	// Indicator that the account opening/maintenance/closing process needs to be treated urgently, that is, sooner than the terms established by the service level agreed between the account holder customer and the account servicing institution.
	UrgencyFlag *YesNoIndicator `xml:"UrgcyFlg,omitempty"`
}

func (a *AccountContract2) SetTargetGoLiveDate(value string) {
	a.TargetGoLiveDate = (*ISODate)(&value)
}

func (a *AccountContract2) SetTargetClosingDate(value string) {
	a.TargetClosingDate = (*ISODate)(&value)
}

func (a *AccountContract2) SetUrgencyFlag(value string) {
	a.UrgencyFlag = (*YesNoIndicator)(&value)
}
