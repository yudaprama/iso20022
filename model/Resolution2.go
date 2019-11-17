package model

// Specifies an item in the agenda of the meeting. Some resolutions are submitted to the vote of the security holders, some are presented for information only.
type Resolution2 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Free text description of the resolution.
	Description *Max1025Text `xml:"Desc,omitempty"`

	// Abbreviated description of the resolution.
	Title *Max350Text `xml:"Titl,omitempty"`

	// Indicates whether a resolution is ordinary, extraordinary or special.
	Type *ResolutionType1Code `xml:"Tp"`

	// Indicates whether the resolution is listed for information or for voting.
	ForInformationOnly *YesNoIndicator `xml:"ForInfOnly"`

	// Indicates whether the resolution is active or withdrawn.
	Status *ResolutionStatus1Code `xml:"Sts"`

	// Indicates whether the resolution has been submitted by the security holder.
	SubmittedBySecurityHolder *YesNoIndicator `xml:"SubmittdBySctyHldr"`

	// Vote options allowed at the resolution level. When specified, it supersedes the vote options given for the meeting.
	VoteInstructionType []*VoteInstruction2Code `xml:"VoteInstrTp,omitempty"`

	// Indicates how the management of the issuing company wishes the security holders to vote.
	ManagementRecommendation *VoteInstruction1Code `xml:"MgmtRcmmndtn,omitempty"`

	// Indicates how the notifying party recommends that the security holders vote.
	NotifyingPartyRecommendation *VoteInstruction1Code `xml:"NtifngPtyRcmmndtn,omitempty"`
}

func (r *Resolution2) SetIssuerLabel(value string) {
	r.IssuerLabel = (*Max35Text)(&value)
}

func (r *Resolution2) SetDescription(value string) {
	r.Description = (*Max1025Text)(&value)
}

func (r *Resolution2) SetTitle(value string) {
	r.Title = (*Max350Text)(&value)
}

func (r *Resolution2) SetType(value string) {
	r.Type = (*ResolutionType1Code)(&value)
}

func (r *Resolution2) SetForInformationOnly(value string) {
	r.ForInformationOnly = (*YesNoIndicator)(&value)
}

func (r *Resolution2) SetStatus(value string) {
	r.Status = (*ResolutionStatus1Code)(&value)
}

func (r *Resolution2) SetSubmittedBySecurityHolder(value string) {
	r.SubmittedBySecurityHolder = (*YesNoIndicator)(&value)
}

func (r *Resolution2) AddVoteInstructionType(value string) {
	r.VoteInstructionType = append(r.VoteInstructionType, (*VoteInstruction2Code)(&value))
}

func (r *Resolution2) SetManagementRecommendation(value string) {
	r.ManagementRecommendation = (*VoteInstruction1Code)(&value)
}

func (r *Resolution2) SetNotifyingPartyRecommendation(value string) {
	r.NotifyingPartyRecommendation = (*VoteInstruction1Code)(&value)
}
