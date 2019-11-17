package model

// Information describing how the voting process is organised.
type VoteParameters1 struct {

	// Number of holdings required for a vote.
	SecuritiesQuantityRequiredToVote *DecimalNumber `xml:"SctiesQtyReqrdToVote,omitempty"`

	// Specifies whether it is allowed to only vote on a part of the entire holding, leaving part of the position un-voted.
	PartialVoteAllowed *YesNoIndicator `xml:"PrtlVoteAllwd"`

	// Specifies whether it is allowed to vote in different directions for the entire holding.
	SplitVoteAllowed *YesNoIndicator `xml:"SpltVoteAllwd"`

	// Date and time by which the vote instructions should be submitted to the intermediary.
	VoteDeadline *DateFormat2Choice `xml:"VoteDdln,omitempty"`

	// Date and time by which the vote instructions should be submitted to the intermediary (STP mode).
	VoteSTPDeadline *DateFormat2Choice `xml:"VoteSTPDdln,omitempty"`

	// Date and time by which the vote instructions should be submitted to the issuer.
	VoteMarketDeadline *DateFormat2Choice `xml:"VoteMktDdln,omitempty"`

	// Indicates the different methods that can be used to vote.
	VoteMethods *VoteMethods `xml:"VoteMthds,omitempty"`

	// Electronic location, e-mail or URL address, where the voting ballot can be requested.
	VotingBallotElectronicAddress *CommunicationAddress4 `xml:"VtngBlltElctrncAdr,omitempty"`

	// Specifies the postal address where the voting ballot can be requested.
	VotingBallotRequestAddress *PostalAddress1 `xml:"VtngBlltReqAdr,omitempty"`

	// Date till which the instructing party can revoke, change or withdraw its voting instruction. This deadline is specified by an intermediary.
	RevocabilityDeadline *DateFormat2Choice `xml:"RvcbltyDdln,omitempty"`

	// Date till which the instructing party can revoke, change or withdraw its voting instruction. This deadline is specified by an intermediary (STP mode).
	RevocabilitySTPDeadline *DateFormat2Choice `xml:"RvcbltySTPDdln,omitempty"`

	// Date till which the instructing party can revoke, change or withdraw its voting instruction. This deadline is set by the issuer.
	RevocabilityMarketDeadline *DateFormat2Choice `xml:"RvcbltyMktDdln,omitempty"`

	// Indicates whether beneficiary details (eg name and address) must be supplied in order to take part to a meeting.
	BeneficialOwnerDisclosure *YesNoIndicator `xml:"BnfclOwnrDsclsr"`

	// Identifies the possible types of voting instructions. When used at the resolution level, it supersedes the value specified in the meeting notice.
	VoteInstructionType []*VoteInstruction2Code `xml:"VoteInstrTp,omitempty"`

	// Cash premium paid to the security holder when voting.
	IncentivePremium *IncentivePremium2 `xml:"IncntivPrm,omitempty"`

	// Date and time by which the vote instructions should be submitted to the intermediary to take advantage of the premium.
	VoteWithPremiumDeadline *DateFormat2Choice `xml:"VoteWthPrmDdln,omitempty"`

	// Date and time by which the vote instructions should be submitted to the intermediary to take advantage of the premium (STP mode).
	VoteWithPremiumSTPDeadline *DateFormat2Choice `xml:"VoteWthPrmSTPDdln,omitempty"`

	// Date and time by which the vote instructions should be submitted to the issuer to take advantage of the premium.
	VoteWithPremiumMarketDeadline *DateFormat2Choice `xml:"VoteWthPrmMktDdln,omitempty"`

	// Additional information on specific requirements for allowing a person to vote.
	AdditionalVotingRequirements *Max350Text `xml:"AddtlVtngRqrmnts,omitempty"`
}

func (v *VoteParameters1) SetSecuritiesQuantityRequiredToVote(value string) {
	v.SecuritiesQuantityRequiredToVote = (*DecimalNumber)(&value)
}

func (v *VoteParameters1) SetPartialVoteAllowed(value string) {
	v.PartialVoteAllowed = (*YesNoIndicator)(&value)
}

func (v *VoteParameters1) SetSplitVoteAllowed(value string) {
	v.SplitVoteAllowed = (*YesNoIndicator)(&value)
}

func (v *VoteParameters1) AddVoteDeadline() *DateFormat2Choice {
	v.VoteDeadline = new(DateFormat2Choice)
	return v.VoteDeadline
}

func (v *VoteParameters1) AddVoteSTPDeadline() *DateFormat2Choice {
	v.VoteSTPDeadline = new(DateFormat2Choice)
	return v.VoteSTPDeadline
}

func (v *VoteParameters1) AddVoteMarketDeadline() *DateFormat2Choice {
	v.VoteMarketDeadline = new(DateFormat2Choice)
	return v.VoteMarketDeadline
}

func (v *VoteParameters1) AddVoteMethods() *VoteMethods {
	v.VoteMethods = new(VoteMethods)
	return v.VoteMethods
}

func (v *VoteParameters1) AddVotingBallotElectronicAddress() *CommunicationAddress4 {
	v.VotingBallotElectronicAddress = new(CommunicationAddress4)
	return v.VotingBallotElectronicAddress
}

func (v *VoteParameters1) AddVotingBallotRequestAddress() *PostalAddress1 {
	v.VotingBallotRequestAddress = new(PostalAddress1)
	return v.VotingBallotRequestAddress
}

func (v *VoteParameters1) AddRevocabilityDeadline() *DateFormat2Choice {
	v.RevocabilityDeadline = new(DateFormat2Choice)
	return v.RevocabilityDeadline
}

func (v *VoteParameters1) AddRevocabilitySTPDeadline() *DateFormat2Choice {
	v.RevocabilitySTPDeadline = new(DateFormat2Choice)
	return v.RevocabilitySTPDeadline
}

func (v *VoteParameters1) AddRevocabilityMarketDeadline() *DateFormat2Choice {
	v.RevocabilityMarketDeadline = new(DateFormat2Choice)
	return v.RevocabilityMarketDeadline
}

func (v *VoteParameters1) SetBeneficialOwnerDisclosure(value string) {
	v.BeneficialOwnerDisclosure = (*YesNoIndicator)(&value)
}

func (v *VoteParameters1) AddVoteInstructionType(value string) {
	v.VoteInstructionType = append(v.VoteInstructionType, (*VoteInstruction2Code)(&value))
}

func (v *VoteParameters1) AddIncentivePremium() *IncentivePremium2 {
	v.IncentivePremium = new(IncentivePremium2)
	return v.IncentivePremium
}

func (v *VoteParameters1) AddVoteWithPremiumDeadline() *DateFormat2Choice {
	v.VoteWithPremiumDeadline = new(DateFormat2Choice)
	return v.VoteWithPremiumDeadline
}

func (v *VoteParameters1) AddVoteWithPremiumSTPDeadline() *DateFormat2Choice {
	v.VoteWithPremiumSTPDeadline = new(DateFormat2Choice)
	return v.VoteWithPremiumSTPDeadline
}

func (v *VoteParameters1) AddVoteWithPremiumMarketDeadline() *DateFormat2Choice {
	v.VoteWithPremiumMarketDeadline = new(DateFormat2Choice)
	return v.VoteWithPremiumMarketDeadline
}

func (v *VoteParameters1) SetAdditionalVotingRequirements(value string) {
	v.AdditionalVotingRequirements = (*Max350Text)(&value)
}
