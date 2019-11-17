package model

// Specifies an item in the agenda of the meeting. Some resolutions are submitted to the vote of the security holders, some are presented for information only.
type Resolution3 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Free text description of the resolution.
	Description *Max1025Text `xml:"Desc,omitempty"`

	// Abbreviated description of the resolution.
	Title *Max350Text `xml:"Titl,omitempty"`

	// Specifies the type of resolution.
	Type *ResolutionType2Code `xml:"Tp,omitempty"`

	// Indicates whether the resolution is listed for information or for voting.
	ForInformationOnly *YesNoIndicator `xml:"ForInfOnly"`

	// Indicates whether the resolution is active or withdrawn.
	Status *ResolutionStatus1Code `xml:"Sts"`

	// Indicates whether the resolution has been submitted by the security holder.
	SubmittedBySecurityHolder *YesNoIndicator `xml:"SubmittdBySctyHldr,omitempty"`

	// Vote options allowed at the resolution level. When specified, it supersedes the vote options given for the meeting.
	VoteInstructionType []*VoteInstruction2Code `xml:"VoteInstrTp,omitempty"`

	// Specifies how the management of the issuing company wishes the security holders to vote.
	ManagementRecommendation *VoteInstruction1Code `xml:"MgmtRcmmndtn,omitempty"`

	// Indicates how the notifying party recommends that the security holders vote.
	NotifyingPartyRecommendation *VoteInstruction1Code `xml:"NtifngPtyRcmmndtn,omitempty"`

	// Number of votes assigned per resolution to one security.
	Entitlement *Entitlement1Choice `xml:"Entitlmnt,omitempty"`
}

func (r *Resolution3) SetIssuerLabel(value string) {
	r.IssuerLabel = (*Max35Text)(&value)
}

func (r *Resolution3) SetDescription(value string) {
	r.Description = (*Max1025Text)(&value)
}

func (r *Resolution3) SetTitle(value string) {
	r.Title = (*Max350Text)(&value)
}

func (r *Resolution3) SetType(value string) {
	r.Type = (*ResolutionType2Code)(&value)
}

func (r *Resolution3) SetForInformationOnly(value string) {
	r.ForInformationOnly = (*YesNoIndicator)(&value)
}

func (r *Resolution3) SetStatus(value string) {
	r.Status = (*ResolutionStatus1Code)(&value)
}

func (r *Resolution3) SetSubmittedBySecurityHolder(value string) {
	r.SubmittedBySecurityHolder = (*YesNoIndicator)(&value)
}

func (r *Resolution3) AddVoteInstructionType(value string) {
	r.VoteInstructionType = append(r.VoteInstructionType, (*VoteInstruction2Code)(&value))
}

func (r *Resolution3) SetManagementRecommendation(value string) {
	r.ManagementRecommendation = (*VoteInstruction1Code)(&value)
}

func (r *Resolution3) SetNotifyingPartyRecommendation(value string) {
	r.NotifyingPartyRecommendation = (*VoteInstruction1Code)(&value)
}

func (r *Resolution3) AddEntitlement() *Entitlement1Choice {
	r.Entitlement = new(Entitlement1Choice)
	return r.Entitlement
}
