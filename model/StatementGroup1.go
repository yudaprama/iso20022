package model

// Group of the statement header reporting the bank services billing and the billing statement.
type StatementGroup1 struct {

	// Identification of a group of customer billing statements.
	GroupIdentification *Max35Text `xml:"GrpId"`

	// Originating financial institution sending the statement.
	Sender *PartyIdentification58 `xml:"Sndr"`

	// Specifies the individual to contact in case of technical problems at the sender's location.
	SenderIndividualContact []*ContactDetails3 `xml:"SndrIndvCtct,omitempty"`

	// Financial institution customer receiving the statement.
	Receiver *PartyIdentification58 `xml:"Rcvr"`

	// Specifies the individual to contact in case of technical problems at the receiver's location.
	ReceiverIndividualContact []*ContactDetails3 `xml:"RcvrIndvCtct,omitempty"`

	// Provides the bank services billing statement recounting of all service chargeable events that occurred during a reporting cycle, such as the end of the month reporting.
	BillingStatement []*BillingStatement1 `xml:"BllgStmt"`
}

func (s *StatementGroup1) SetGroupIdentification(value string) {
	s.GroupIdentification = (*Max35Text)(&value)
}

func (s *StatementGroup1) AddSender() *PartyIdentification58 {
	s.Sender = new(PartyIdentification58)
	return s.Sender
}

func (s *StatementGroup1) AddSenderIndividualContact() *ContactDetails3 {
	newValue := new(ContactDetails3)
	s.SenderIndividualContact = append(s.SenderIndividualContact, newValue)
	return newValue
}

func (s *StatementGroup1) AddReceiver() *PartyIdentification58 {
	s.Receiver = new(PartyIdentification58)
	return s.Receiver
}

func (s *StatementGroup1) AddReceiverIndividualContact() *ContactDetails3 {
	newValue := new(ContactDetails3)
	s.ReceiverIndividualContact = append(s.ReceiverIndividualContact, newValue)
	return newValue
}

func (s *StatementGroup1) AddBillingStatement() *BillingStatement1 {
	newValue := new(BillingStatement1)
	s.BillingStatement = append(s.BillingStatement, newValue)
	return newValue
}
