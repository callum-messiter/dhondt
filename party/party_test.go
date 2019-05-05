package party_test

import (
	"github.com/dhondt/party"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParty_AllocateSeatToCandidate(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name       string
		partyID    party.ID
		votesWon   int
		candidates []string
		expected   []*party.Party
	}{
		{
			name:       "allocate seats to candidates, in order",
			partyID:    party.Cons,
			votesWon:   100,
			candidates: []string{"A", "B", "C"},
			expected: []*party.Party{
				{
					ID:                party.Cons,
					VotesWon:          100,
					Candidates:        []string{"A", "B", "C"},
					SeatsWon:          1,
					CandidatesElected: []string{"A"},
					Quotient:          50,
				},
				{
					ID:                party.Cons,
					VotesWon:          100,
					Candidates:        []string{"A", "B", "C"},
					SeatsWon:          2,
					CandidatesElected: []string{"A", "B"},
					Quotient:          33.33,
				},
				{
					ID:                party.Cons,
					VotesWon:          100,
					Candidates:        []string{"A", "B", "C"},
					SeatsWon:          3,
					CandidatesElected: []string{"A", "B", "C"},
					Quotient:          25,
				},
			},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := party.New(tc.partyID, tc.votesWon, tc.candidates)
			for i := range tc.candidates {
				got.AllocateSeatToCandidate()
				assert.Equal(t, tc.expected[i], got)
			}
		})
	}
}

func TestParties_HighestQuotient(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		parties  party.Parties
		expected *party.Party
	}{
		{
			name: "labour has highest quotient by 0.01 margin",
			parties: party.Parties{
				party.Cons:   &party.Party{ID: party.Cons, Quotient: 3.38},
				party.Lab:    &party.Party{ID: party.Lab, Quotient: 3.39},
				party.LibDem: &party.Party{ID: party.Lab, Quotient: 3.37},
			},
			expected: &party.Party{ID: party.Lab, Quotient: 3.39},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.parties.RoundWinner()

			assert.Equal(t, tc.expected, got)
		})
	}
}
