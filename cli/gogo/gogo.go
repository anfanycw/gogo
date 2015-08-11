package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

type Vacancy struct {
	Rooms []string `json:"vacant"`
}

func main() {
	app := cli.NewApp()
	app.Name = "gogo"
	app.Usage = "Find out which restroom is open when you've gotta go"
	app.Action = func(c *cli.Context) {
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

	app.Run(os.Args)
}
