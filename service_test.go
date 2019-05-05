package dhondt_test

import (
	"github.com/dhondt"
	"github.com/dhondt/party"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_AllocateSeats(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		input    party.Parties
		seats    int
		expected dhondt.ElectionResult
	}{
		{
			name:  "2004 EU elections - South East",
			seats: 10,
			input: party.Parties{
				party.Cons: &party.Party{
					ID:       party.Cons,
					VotesWon: 776370,
					Candidates: []string{
						"Daniel Hannan",
						"Nirj Deva",
						"James Elles",
						"Richard Ashworth",
						"Roy Perry",
						"Therese Coffey",
						"David Logan",
						"Ferris Cowper",
						"Richard Robinson",
					},
				},
				party.Green: &party.Party{
					ID:       party.Green,
					VotesWon: 173351,
					Candidates: []string{
						"Caroline Lucas",
						"Mike Woodin",
						"Miriam Kennet",
						"Keith Taylor",
						"Alan Francis",
						"Xanthe Bevis",
						"Hazel Dawe",
						"Derek Wall",
						"nthony Cooper",
						"Michael Stimson",
					},
				},
				party.Lab: &party.Party{
					ID:       party.Lab,
					VotesWon: 301398,
					Candidates: []string{
						"Peter Skinner",
						"Mark Watts",
						"Ann Davison",
						"Simon Burgess",
						"Janet Sully",
						"Mark Muller",
						"Josephine Wood",
						"Raj Chandarana",
						"Gillian Roles",
						"David Menon",
					},
				},
				party.LibDem: &party.Party{
					ID:       party.LibDem,
					VotesWon: 338342,
					Candidates: []string{
						"Chris Huhne",
						"Baroness Nicholson of Winterbourne",
						"Sharon Bowles",
						"Catherine Bearder",
						"James Walsh",
						"Ann Lee",
						"John Vincent",
						"John Ford",
						"Charles Fraser-Fleming",
						"James Barnard",
					},
				},
				party.UKIP: &party.Party{
					ID:       party.UKIP,
					VotesWon: 431111,
					Candidates: []string{
						"Nigel Farage",
						"Ashley Mote",
						"David Lott",
						"Craig Mackinlay",
						"Timothy Cross",
						"Petrina Holdsworth",
						"David Abbott",
						"Stephen Harris",
						"Michael Wigley",
						"Lisa Hawkins",
					},
				},
				party.BNP: &party.Party{
					ID:       party.BNP,
					VotesWon: 64877,
					Candidates: []string{
						"Brian Galloway",
						"Julie Russell",
						"Timothy Rait",
						"Peter Lane",
						"Roger Robertson",
						"Julian Crewe",
						"Adam Champneys",
						"Ian Johnson",
						"Dennis Whiting",
						"Vernon Atkinson",
					},
				},
				party.SCP: &party.Party{
					ID:       party.SCP,
					VotesWon: 42861,
					Candidates: []string{
						"Grahame Leon-Smith",
						"David Gray",
						"Patrick Eston",
						"Rona Brown",
						"Paresh Kotecha",
						"Larry Kreeger",
						"Michael Devine",
						"Terry Patinson",
						"Ian Murdoch",
						"Alfred Egleton",
					},
				},
				party.EngDem: &party.Party{
					ID:       party.EngDem,
					VotesWon: 29126,
					Candidates: []string{
						"Steven Uncles",
						"Robert Sulley",
						"Courtney Williams",
						"Richard Sutton",
						"Jacqueline Brookman",
						"David Uncles",
						"Louise Uncles",
					},
				},
				party.Respect: &party.Party{
					ID:       party.Respect,
					VotesWon: 13426,
					Candidates: []string{
						"Patrick O'Keeffe",
						"Ingrid Dodd",
						"Muriel Hirsch",
						"Ajaz Khan",
						"Sally Watkins",
						"Jonathan Molyneux",
						"Norman Thomas",
						"Ella Noyes",
						"Bunny La Roche",
						"Angelina Rai",
					},
				},
				party.Peace: &party.Party{
					ID:       party.Peace,
					VotesWon: 12572,
					Candidates: []string{
						"John Morris",
						"Caroline O'Reilly",
						"Geoffrey Pay",
						"Rachel Hancock",
						"James Duggan",
						"Kate Hebden",
						"Cyril Bolam",
						"Carol Morris",
						"Anne Brewer",
					},
				},
				party.ChristianPeoples: &party.Party{
					ID:       party.ChristianPeoples,
					VotesWon: 11733,
					Candidates: []string{
						"David John Bamber",
						"David Campanale",
						"Gladstone Macaulay",
					},
				},
				party.ProLifeAlliance: &party.Party{
					ID:       party.ProLifeAlliance,
					VotesWon: 6579,
					Candidates: []string{
						"Dominica Roberts",
						"Gillian Duval",
						"Josephine Quintavalle",
						"Penelope Orford",
						"Mark Carroll",
						"Rebecca Ng",
						"John Dixon",
						"Francis O'Brien",
						"Yvonne Windsor",
						"Carl St John",
					},
				},
				party.Independent: &party.Party{
					ID:       party.Independent,
					VotesWon: 5671,
					Candidates: []string{
						"Philip Rhodes",
					},
				},
			},
			expected: dhondt.ElectionResult{
				party.Cons: {
					TotalWon:          4,
					CandidatesElected: []string{"Daniel Hannan", "Nirj Deva", "James Elles", "Richard Ashworth"},
				},
				party.Green: {
					TotalWon:          1,
					CandidatesElected: []string{"Caroline Lucas"},
				},
				party.Lab: {
					TotalWon:          1,
					CandidatesElected: []string{"Peter Skinner"},
				},
				party.LibDem: {
					TotalWon:          2,
					CandidatesElected: []string{"Chris Huhne", "Baroness Nicholson of Winterbourne"},
				},
				party.UKIP: {
					TotalWon:          2,
					CandidatesElected: []string{"Nigel Farage", "Ashley Mote"},
				},
				party.BNP:              {TotalWon: 0, CandidatesElected: nil},
				party.SCP:              {TotalWon: 0, CandidatesElected: nil},
				party.EngDem:           {TotalWon: 0, CandidatesElected: nil},
				party.Respect:          {TotalWon: 0, CandidatesElected: nil},
				party.Peace:            {TotalWon: 0, CandidatesElected: nil},
				party.ChristianPeoples: {TotalWon: 0, CandidatesElected: nil},
				party.ProLifeAlliance:  {TotalWon: 0, CandidatesElected: nil},
				party.Independent:      {TotalWon: 0, CandidatesElected: nil},
			},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := dhondt.Service{AvailableSeats: tc.seats}
			got := s.AllocateSeats(parties(tc.input))
			assert.Equal(t, tc.expected, got)
		})
	}
}

func parties(partyData party.Parties) party.Parties {
	parties := make(party.Parties, len(partyData))
	for _, p := range partyData {
		parties[p.ID] = party.New(p.ID, p.VotesWon, p.Candidates)
	}
	return parties
}
