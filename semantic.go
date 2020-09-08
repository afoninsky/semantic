package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/afoninsky/semantic/pkg/replace"
	"github.com/afoninsky/semantic/pkg/repository"
	"github.com/afoninsky/semantic/pkg/static"
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

	// return git semantic helpers, usage example:
	// # source <(semantic aliases)
	// # gfeat added new semantic feature
	case "aliases":
		data, err := static.Asset("scripts/aliases.sh")
		exitIfErr(err)
		fmt.Println(string(data))

	// replace patterns in files, used as replacement for tools like "sed"
	// # replace test.yaml "version: (.+)" "version: new"
	// # replace test.yaml "version: (.+)" "version: $1-release"
	case "replace":
		if len(os.Args) < 5 {
			exitIfErr(errors.New("usage: replace <file> <pattern> <value>"))
		}
		fmt.Printf("%s: \"%s\" -> \"%s\"\n", os.Args[2], os.Args[3], os.Args[4])
		exitIfErr(replace.Do(os.Args[2], os.Args[3], os.Args[4]))

	// return current release version, useful in CI
	case "current":
		fmt.Printf(info.LatestVersion)

	// return next release version, useful in CI
	case "next":
		fmt.Printf(info.NextVersion)

	// return current tag containing version and git commit, useful in CI
	case "tag":
		fmt.Printf(info.CurrentTag)

	// display common information about status of semantic release
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
