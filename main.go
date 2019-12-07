package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"
)

func main() {

	competitions := []int{
		2014, // Liga
		2015, // Ligue 1
		2002, // Bundesliga
		2019, // Serie A
		2021, // Premier League
	}

	var ranking []team

	for _, c := range competitions {

		println(fmt.Sprintf("Getting competition : %d", c))

		s, err := getStanding(c)

		if err != nil {
			log.Fatal(err)
		}

		ranking = append(ranking, formatStanding(s)...)

		//time.Sleep(5 * time.Second)
	}

	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].AveragePoints > ranking[j].AveragePoints
	})

	tmpl, err := template.ParseFiles("web/layout.html")

	if err != nil {
		println("Could not parse layout.html")
		log.Fatal(err)
	}

	println("Parsed layout.html with success")

	file, err := os.Create("build/index.html")
	defer file.Close()

	if err != nil {
		println("Could not create build/index.html")
		log.Fatal(err)
	}

	err = tmpl.Execute(file, struct{ Ranking []team }{ranking})

	if err != nil {
		println("Could not populate build/index.html")
		log.Fatal(err)
	}

	println("Site built with success")
}

func formatStanding(s *standing) (teams []team) {
	for _, row := range s.Standings[0].Table {
		teams = append(teams, team{
			Name:          row.Team.Name,
			Points:        row.Points,
			AveragePoints: float32(row.Points) / float32(row.PlayedGames),
		})
	}

	return
}

func getStanding(competitionID int) (s *standing, err error) {
	b := request(fmt.Sprintf("https://api.football-data.org/v2/competitions/%d/standings?standingType=TOTAL", competitionID))

	err = json.Unmarshal(b, &s)

	return
}
