package model

// Specifies linkage information of a corporate action message.
type LinkedCorporateAction1 struct {

	// The function of the notification e.g. new notification.
	NotificationType *CorporateActionNotificationType1Code `xml:"NtfctnTp"`

	// The identification of the linked notification advice.
	LinkedAgentCANotificationAdviceIdentification *DocumentIdentification8 `xml:"LkdAgtCANtfctnAdvcId,omitempty"`

	// Specifies when the instruction is to be executed relative to a linked instruction.
	LinkageType *ProcessingPosition2FormatChoice `xml:"LkgTp,omitempty"`

	// Reference given to the linked event by the CA event issuer (agent).
	LinkedIssuerCorporateActionIdentification *Max35Text `xml:"LkdIssrCorpActnId,omitempty"`

	// Reference assigned by the CSD to the linked coporate avent.
	LinkedCorporateActionProcessingIdentification *Max35Text `xml:"LkdCorpActnPrcgId,omitempty"`
}

func (l *LinkedCorporateAction1) SetNotificationType(value string) {
	l.NotificationType = (*CorporateActionNotificationType1Code)(&value)
}

func (l *LinkedCorporateAction1) AddLinkedAgentCANotificationAdviceIdentification() *DocumentIdentification8 {
	l.LinkedAgentCANotificationAdviceIdentification = new(DocumentIdentification8)
	return l.LinkedAgentCANotificationAdviceIdentification
}

func (l *LinkedCorporateAction1) AddLinkageType() *ProcessingPosition2FormatChoice {
	l.LinkageType = new(ProcessingPosition2FormatChoice)
	return l.LinkageType
}

func (l *LinkedCorporateAction1) SetLinkedIssuerCorporateActionIdentification(value string) {
	l.LinkedIssuerCorporateActionIdentification = (*Max35Text)(&value)
}

func (l *LinkedCorporateAction1) SetLinkedCorporateActionProcessingIdentification(value string) {
	l.LinkedCorporateActionProcessingIdentification = (*Max35Text)(&value)
}
