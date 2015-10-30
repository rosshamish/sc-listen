package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
	"net/url"
	"os"
)

func appAction(c *cli.Context) {
	var scUser string
	scUser = "rosshamish"
	if len(c.Args()) > 0 {
		scUser = c.Args()[0]
	}

	scapi := &soundcloud.Api{
		ClientId: "73275ddb505720361c57d1f3b3e837c0",
	}

	users, err := scapi.Users(url.Values{"q": []string{scUser}})
	if err != nil {
		fmt.Println(err)
		return
	} else if len(users) == 0 {
		fmt.Println("No users found. Exiting.")
		return
	} else if len(users) > 1 {
		fmt.Println("Did you mean...")
		for i, user := range users {
			if i >= 3 {
				break
			}
			fmt.Println("\t%s", user.Username)
		}
		fmt.Println()
	}

	user := users[0]
	fmt.Printf("Found user %s\n", user.Username)
	fmt.Printf("\tFull Name: %s\n", user.FullName)
	fmt.Printf("\tCountry: %s\n", user.Country)

	uapi := scapi.User(user.Id)
	tracks, err := uapi.Tracks(url.Values{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tracks:")
	for i, track := range tracks {
		if i >= 5 {
			fmt.Printf("\t...\n")
			break
		}
		fmt.Printf("\t%s: %s\n", user.Username, track.Title)
		//fmt.Println("\t\t%s\n", track.StreamUrl)
	}

	playlists, err := uapi.Playlists(url.Values{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Playlists:")
	for i, playlist := range playlists {
		if i >= 5 {
			fmt.Println("\t...")
			break
		}
		fmt.Printf("\t%s: %s\n", user.Username, playlist.Title)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "sc-listen"
	app.Usage = "stream soundcloud from the terminal"
	app.Action = appAction

	app.Run(os.Args)
}
