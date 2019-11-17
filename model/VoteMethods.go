package model

// List of the different methods that can be used to vote.
type VoteMethods struct {

	// Network address through which a voting party can cast its vote via a structured message.
	VoteThroughNetwork *AnyBICIdentifier `xml:"VoteThrghNtwk,omitempty"`

	// Specifies the address where voting ballot can be sent.
	VoteByMail *PostalAddress1 `xml:"VoteByMail,omitempty"`

	// Electronic address, e-mail or website, where a security holder can vote.
	ElectronicVote *CommunicationAddress4 `xml:"ElctrncVote,omitempty"`

	// Telephone number providing access to an automated voting system.
	VoteByTelephone *Max35Text `xml:"VoteByTel,omitempty"`
}

func (v *VoteMethods) SetVoteThroughNetwork(value string) {
	v.VoteThroughNetwork = (*AnyBICIdentifier)(&value)
}

func (v *VoteMethods) AddVoteByMail() *PostalAddress1 {
	v.VoteByMail = new(PostalAddress1)
	return v.VoteByMail
}

func (v *VoteMethods) AddElectronicVote() *CommunicationAddress4 {
	v.ElectronicVote = new(CommunicationAddress4)
	return v.ElectronicVote
}

func (v *VoteMethods) SetVoteByTelephone(value string) {
	v.VoteByTelephone = (*Max35Text)(&value)
}
