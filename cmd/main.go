package main

import (
	"fmt"
	"os"

	"github.com/afoninsky/semantic/pkg/repository"
)

func main() {
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	r, err := repository.New("./")
	exitIfErr(err)
	info, err := r.Info()
	exitIfErr(err)

	switch cmd {

	case "current":
		fmt.Printf(info.LatestVersion)

	case "next":
		fmt.Printf(info.NextVersion)

	case "tag":
		fmt.Printf(info.CurrentTag)

	case "":
		fmt.Printf("Latest version: %s\n", info.LatestVersion)
		fmt.Printf("Current tag: %s\n", info.CurrentTag)
		if info.NextVersion != "" {
			fmt.Println("Commits since latest version:")
			for _, c := range info.NextCommits {
				if c.Type != "" {
					fmt.Printf("\t * %s: %s\n", c.Type, c.Message)
				}
			}
			fmt.Printf("Next possible version: %s\n", info.NextVersion)
		}

	default:
		exitIfErr(fmt.Errorf("invalid command: %s", cmd))
	}

}

func exitIfErr(err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}
}
