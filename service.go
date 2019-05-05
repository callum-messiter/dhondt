package dhondt

import (
	"github.com/dhondt/party"
)

// Service is used to generate election results based on the D'Hondt method of party-list proportional representation.
type Service struct {
	AvailableSeats int
}

// AllocateSeats implements the D'Hondt method of party-list proportional representation.
// Each seat is allocated to a candidate from whichever party wins that round. The winner of the first round is the
// party with the most votes; the party with the highest quotient wins each round thereafter.
func (s *Service) AllocateSeats(parties party.Parties) ElectionResult {
	for i := 0; i < s.AvailableSeats; i++ {
		parties.RoundWinner().AllocateSeatToCandidate()
	}

	result := make(ElectionResult, len(parties))
	for p, data := range parties {
		result[p] = party.SeatAllocation{
			TotalWon:          data.SeatsWon,
			CandidatesElected: data.CandidatesElected,
		}
	}
	return result
}

// ElectionResult contains proportional seat allocation by party, in an election.
type ElectionResult map[party.ID]party.SeatAllocation
