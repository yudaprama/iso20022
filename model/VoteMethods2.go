package model

// List of the different methods that can be used to vote.
type VoteMethods2 struct {

	// Network address through which a voting party can cast its vote via a structured message.
	VoteThroughNetwork []*AnyBICIdentifier `xml:"VoteThrghNtwk,omitempty"`

	// Specifies the address where voting ballot can be sent.
	VoteByMail []*PostalAddress1 `xml:"VoteByMail,omitempty"`

	// Electronic address, e-mail or website, where a security holder can vote.
	ElectronicVote []*CommunicationAddress4 `xml:"ElctrncVote,omitempty"`

	// Telephone number providing access to an automated voting system.
	VoteByTelephone []*Max35Text `xml:"VoteByTel,omitempty"`
}

func (v *VoteMethods2) AddVoteThroughNetwork(value string) {
	v.VoteThroughNetwork = append(v.VoteThroughNetwork, (*AnyBICIdentifier)(&value))
}

func (v *VoteMethods2) AddVoteByMail() *PostalAddress1 {
	newValue := new(PostalAddress1)
	v.VoteByMail = append(v.VoteByMail, newValue)
	return newValue
}

func (v *VoteMethods2) AddElectronicVote() *CommunicationAddress4 {
	newValue := new(CommunicationAddress4)
	v.ElectronicVote = append(v.ElectronicVote, newValue)
	return newValue
}

func (v *VoteMethods2) AddVoteByTelephone(value string) {
	v.VoteByTelephone = append(v.VoteByTelephone, (*Max35Text)(&value))
}
