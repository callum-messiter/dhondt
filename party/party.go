package party

import (
	"math"
)

// Party is used to store dynamic, round-by-round election results for a particular political party.
type Party struct {
	ID                ID
	VotesWon          int
	Candidates        []string
	Quotient          float64
	SeatsWon          int
	CandidatesElected []string
}

// New initialises a Party, setting the starting quotient as the number of votes won.
func New(id ID, votesWon int, candidates []string) *Party {
	return &Party{
		ID:         id,
		VotesWon:   votesWon,
		Candidates: candidates,
		Quotient:   float64(votesWon),
		SeatsWon:   0,
	}
}

// ID is a unique reference to party.
type ID int

const (
	Cons ID = iota
	Green
	Lab
	LibDem
	UKIP
	BNP
	SCP
	EngDem
	Respect
	Peace
	ChristianPeoples
	ProLifeAlliance
	Independent
)

// AllocateSeatToCandidate assigns the party's won seat to a candidate from the party's list. The party selects the
// order of election of their candidates. E.g. If their list is [PersonA, PersonB], the first seat the party
// wins will be allocated to PersonA; if they win a second seat, it will be allocated to PersonB.
func (d *Party) AllocateSeatToCandidate() {
	d.CandidatesElected = append(d.CandidatesElected, d.Candidates[d.SeatsWon])
	d.SeatsWon++
	quotient := float64(d.VotesWon) / float64(d.SeatsWon+1)
	d.Quotient = math.Round(quotient*100) / 100
}

// Parties stores party election data that are referenced by the party's unique id.
type Parties map[ID]*Party

// RoundWinner determines, based on the current quotients of the parties, which party wins the seat.
func (p *Parties) RoundWinner() *Party {
	partyWithHighestQuotient := &Party{Quotient: 0}
	for _, party := range *p {
		if party.Quotient > partyWithHighestQuotient.Quotient {
			partyWithHighestQuotient = party
		}
	}
	return partyWithHighestQuotient
}

// SeatAllocation contains data about the party's won seats and elected candidates.
type SeatAllocation struct {
	TotalWon          int
	CandidatesElected []string
}
