package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name          string
	matchesPlayed int
	wins          int
	losses        int
	draws         int
}

type teams map[string]*team

func (t *team) points() int {
	return 3*t.wins + 1*t.draws
}

func (t teams) sort() []*team {
	var teams []*team
	for _, team := range t {
		teams = append(teams, team)
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points() == teams[j].points() {
			return teams[i].name < teams[j].name
		}
		return teams[i].points() > teams[j].points()
	})
	return teams
}

func printResults(writer io.Writer, teams teams) (err error) {
	_, err = fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}

	for _, team := range teams.sort() {
		_, err = fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name, team.matchesPlayed, team.wins, team.draws, team.losses, team.points())
		if err != nil {
			return err
		}
	}

	return err
}

func Tally(reader io.Reader, writer io.Writer) (err error) {
	teams := make(teams)
	s := bufio.NewScanner(reader)
	for s.Scan() {
		l := s.Text()
		if l == "" || l[0] == '#' {
			continue
		}
		chunks := strings.Split(l, ";")
		if len(chunks) != 3 {
			return fmt.Errorf("invalid score entry")
		}
		team1, team2, outcome := chunks[0], chunks[1], chunks[2]
		if teams[team1] == nil {
			teams[team1] = &team{name: team1}
		}
		if teams[team2] == nil {
			teams[team2] = &team{name: team2}
		}
		switch outcome {
		case "win":
			teams[team1].matchesPlayed++
			teams[team2].matchesPlayed++
			teams[team1].wins++
			teams[team2].losses++
		case "loss":
			teams[team1].matchesPlayed++
			teams[team2].matchesPlayed++
			teams[team1].losses++
			teams[team2].wins++
		case "draw":
			teams[team1].matchesPlayed++
			teams[team2].matchesPlayed++
			teams[team1].draws++
			teams[team2].draws++
		default:
			return fmt.Errorf("invalid outcome %s", outcome)
		}

	}
	err = printResults(writer, teams)
	return err
}
