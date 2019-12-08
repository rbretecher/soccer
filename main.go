package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"
)

func main() {

	competitions := []int{
		2001, // UEFA Champions League
		2014, // Liga
		2015, // Ligue 1
		2002, // Bundesliga
		2019, // Serie A
		2021, // Premier League
	}

	mRanking := map[int]*team{}

	for _, c := range competitions {

		println(fmt.Sprintf("Getting competition : %d", c))

		s, err := getStanding(c)

		if err != nil {
			log.Fatal(err)
		}

		addToRanking(mRanking, s)

		// Token only allows 10 requests/min
		//time.Sleep(2 * time.Second)
	}

	ranking := make([]*team, 0)
	for _, team := range mRanking {
		ranking = append(ranking, team)
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

	if _, err := os.Stat("build"); os.IsNotExist(err) {
		err = os.Mkdir("build", 0700)

		if err != nil {
			println("Could not create build folder")
			log.Fatal(err)
		}
	}

	file, err := os.Create("build/index.html")
	defer file.Close()

	if err != nil {
		println("Could not create build/index.html")
		log.Fatal(err)
	}

	err = tmpl.Execute(file, struct{ Ranking []*team }{ranking})

	if err != nil {
		println("Could not populate build/index.html")
		log.Fatal(err)
	}

	println("Site built with success")
}

func addToRanking(ranking map[int]*team, s *standing) {
	for _, currentStanding := range s.Standings {
		for _, row := range currentStanding.Table {
			if _, ok := ranking[row.Team.ID]; ok {
				ranking[row.Team.ID].Points += row.Points
				ranking[row.Team.ID].PlayedGames += row.PlayedGames
				ranking[row.Team.ID].AveragePoints = float32(ranking[row.Team.ID].Points) / float32(ranking[row.Team.ID].PlayedGames)
			} else {
				ranking[row.Team.ID] = &team{
					Name:          row.Team.Name,
					Points:        row.Points,
					PlayedGames:   row.PlayedGames,
					AveragePoints: float32(row.Points) / float32(row.PlayedGames),
				}
			}
		}
	}
}
