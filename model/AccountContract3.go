package model

// Specifies target and actual dates dates related to account opening and closing.
type AccountContract3 struct {

	// Date on which the account and related basic services  are expected to be operational for the account owner.
	TargetGoLiveDate *ISODate `xml:"TrgtGoLiveDt,omitempty"`

	// Date on which the account and related services are expected to cease to be operational for the account owner.
	TargetClosingDate *ISODate `xml:"TrgtClsgDt,omitempty"`

	// Date on which the account and related basic services are effectively operational for the account owner.
	GoLiveDate *ISODate `xml:"GoLiveDt,omitempty"`

	// Date on which the account and related services cease effectively to be operational for the account owner.
	ClosingDate *ISODate `xml:"ClsgDt,omitempty"`

	// Indicator that the account opening/maintenance/closing process needs to be treated urgently, that is, sooner than the terms established by the service level agreed between the account holder customer and the account servicing institution.
	UrgencyFlag *YesNoIndicator `xml:"UrgcyFlg,omitempty"`

	// Indicates removal of the account. After removal, an account will not appear anymore in reports.
	RemovalIndicator *YesNoIndicator `xml:"RmvlInd,omitempty"`
}

func (a *AccountContract3) SetTargetGoLiveDate(value string) {
	a.TargetGoLiveDate = (*ISODate)(&value)
}

func (a *AccountContract3) SetTargetClosingDate(value string) {
	a.TargetClosingDate = (*ISODate)(&value)
}

func (a *AccountContract3) SetGoLiveDate(value string) {
	a.GoLiveDate = (*ISODate)(&value)
}

func (a *AccountContract3) SetClosingDate(value string) {
	a.ClosingDate = (*ISODate)(&value)
}

func (a *AccountContract3) SetUrgencyFlag(value string) {
	a.UrgencyFlag = (*YesNoIndicator)(&value)
}

func (a *AccountContract3) SetRemovalIndicator(value string) {
	a.RemovalIndicator = (*YesNoIndicator)(&value)
}
