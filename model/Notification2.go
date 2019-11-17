package model

// Information about the type of notification required.
type Notification2 struct {

	// Type of notification.
	NotificationType *Max35Text `xml:"NtfctnTp"`

	// Indicates whether the notification is required.
	Required *YesNoIndicator `xml:"Reqrd"`

	// Specifies how the notification is sent.
	DistributionType *InformationDistribution1Choice `xml:"DstrbtnTp,omitempty"`
}

func (n *Notification2) SetNotificationType(value string) {
	n.NotificationType = (*Max35Text)(&value)
}

func (n *Notification2) SetRequired(value string) {
	n.Required = (*YesNoIndicator)(&value)
}

func (n *Notification2) AddDistributionType() *InformationDistribution1Choice {
	n.DistributionType = new(InformationDistribution1Choice)
	return n.DistributionType
}
