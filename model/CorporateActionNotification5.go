package model

// Corporate action event notification status and contents.
type CorporateActionNotification5 struct {

	// Specifies the type of notification.
	NotificationType *CorporateActionNotificationType1Code `xml:"NtfctnTp"`

	// Specifies the status of the details of the corporate action event.
	ProcessingStatus *CorporateActionProcessingStatus5Choice `xml:"PrcgSts"`

	// Indicates whether the eligible balance is final except for a voluntary corporate action event where it can represent the current eligible balance when communicated before expiration date of that event.
	EligibleBalanceIndicator *YesNoIndicator `xml:"ElgblBalInd,omitempty"`
}

func (c *CorporateActionNotification5) SetNotificationType(value string) {
	c.NotificationType = (*CorporateActionNotificationType1Code)(&value)
}

func (c *CorporateActionNotification5) AddProcessingStatus() *CorporateActionProcessingStatus5Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus5Choice)
	return c.ProcessingStatus
}

func (c *CorporateActionNotification5) SetEligibleBalanceIndicator(value string) {
	c.EligibleBalanceIndicator = (*YesNoIndicator)(&value)
}
