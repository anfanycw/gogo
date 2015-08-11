package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

type Vacancy struct {
	Rooms []string `json:"vacant"`
}

func main() {
	app := cli.NewApp()
	app.Name = "gogo"
	app.Usage = "Find out which restroom is open when you've gotta go"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "reserve, r",
			Value: "",
			Usage: "restroom location",
		},
		cli.StringFlag{
			Name:  "name, n",
			Value: "",
			Usage: "reserver name",
		},
	}
	app.Action = func(c *cli.Context) {
		r := c.String("reserve")
		n := c.String("name")

		if r != "" && n != "" {
			reserveRoom(r, n)
		} else {
			getVacancy()
		}
	}

	app.Run(os.Args)
}

func getVacancy() {
	println("Vacancy")

	resp, err := http.Get("http://ec2-52-27-166-124.us-west-2.compute.amazonaws.com:8080/occupancy")
	if err != nil {
		println(fmt.Sprintf("Sorry, the service is unavailable: %v", err))
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	vacancy := &Vacancy{}
	err = json.Unmarshal(body, vacancy)
	if err != nil {
		println(fmt.Sprintf("Problem reading data: %v", err))
		return
	}

	println("=======")

	if len(vacancy.Rooms) == 0 {
		println("None")
	}

	for index, value := range vacancy.Rooms {
		println(fmt.Sprintf("%d. %s", index+1, value))
	}
}

func reserveRoom(room, name string) {
	println(fmt.Sprintf("reserving restroom: %s by %s", room, name))
	resp, err := http.Post(
		"http://ec2-52-27-166-124.us-west-2.compute.amazonaws.com:8080/reservation",
		"application/json",
		strings.NewReader(fmt.Sprintf(`{"location":"%s", "name":"%s"}`, room, name)),
	)
	if err != nil {
		println(fmt.Sprintf("Problem reserving room: %v", err))
		return
	}

	if resp.StatusCode == 404 {
		println(fmt.Sprintf("Room '%s' not found"))
		return
	}

	if resp.StatusCode == 400 {
		println(fmt.Sprintf("Room is occupied"))
		return
	}

	println(fmt.Sprintf("Room reserved: %v", resp.StatusCode))
}
