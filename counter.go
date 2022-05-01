package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"unicode/utf8"
)

var (
	flagHelp = flag.Bool("h", false, "Shows usage options.")
	flagUser = flag.String("u", "", "The user to be displayed.")
	flagOrg  = flag.String("o", "", "The organization to be displayed.")
)

func main() {
	flag.Parse()
	if *flagHelp {
		fmt.Println("Usage: [options]")
		flag.PrintDefaults()
		return
	}
	if *flagUser == "" && *flagOrg == "" {
		log.Fatalln("Please specify a user or organization.")
	}
	if len(*flagUser) > 0 {
		displayRepos("users", *flagUser)
	} else if len(*flagOrg) > 0 {
		displayRepos("orgs", *flagOrg)
	}
}

func displayRepos(category string, owner string) {
	fmt.Println("Please waiting...")
	repos, size, stars, forks, watchers, maxNameLen := getRepos(category, owner)
	sortRepos(repos)
	nameWidth := 0
	if maxNameLen%10 == 0 {
		nameWidth = maxNameLen + 10
	} else {
		nameWidth = maxNameLen + (10 - maxNameLen%10)
	}
	fmt.Println(strings.Repeat("=", nameWidth+60))
	for _, repo := range repos {
		fmt.Print(repo.Name, strings.Repeat(" ", nameWidth-utf8.RuneCountInString(repo.Name)))
		fmt.Printf("%-16sStars %-10dForks %-10dWatchers %-10d\n", repo.Language, repo.Stars, repo.Forks, repo.Watchers)
	}
	fmt.Println(strings.Repeat("=", nameWidth+60))
	fmt.Print("Total", strings.Repeat(" ", nameWidth-5))
	fmt.Printf("Repos %-10dStars %-10dForks %-10dWatchers %-10d\n", size, stars, forks, watchers)
}

func getRepos(category string, owner string) ([]*Repository, int, int, int, int, int) {
	page := 1
	repos := make([]*Repository, 10)
	size := 0
	stars := 0
	forks := 0
	watchers := 0
	maxNameLen := 0
	for true {
		url := fmt.Sprintf("https://api.github.com/%s/%s/repos?page=%d&per_page=%d", category, owner, page, 100)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return repos[:size], size, stars, forks, watchers, maxNameLen
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println(err)
			return repos[:size], size, stars, forks, watchers, maxNameLen
		}
		var list []*Repository
		err = json.Unmarshal(body, &list)
		if err != nil {
			fmt.Println(err)
			return repos[:size], size, stars, forks, watchers, maxNameLen
		}
		if len(list) == 0 {
			break
		}
		for _, repo := range list {
			if repo.Fork {
				continue
			}
			if len(repos) == size {
				tmp := make([]*Repository, size<<1)
				copy(tmp, repos)
				repos = tmp
			}
			repos[size] = repo
			size++
			stars += repo.Stars
			forks += repo.Forks
			watchers += repo.Watchers
			nameLen := utf8.RuneCountInString(repo.Name)
			if nameLen > maxNameLen {
				maxNameLen = nameLen
			}
		}
		page++
	}
	return repos[:size], size, stars, forks, watchers, maxNameLen
}

func sortRepos(repos []*Repository) {
	sort.Slice(repos, func(i, j int) bool {
		a := repos[i]
		b := repos[j]
		if a.Stars != b.Stars {
			return a.Stars >= b.Stars
		} else if a.Forks != b.Forks {
			return a.Forks >= b.Forks
		} else {
			return strings.ToLower(a.Name) < strings.ToLower(b.Name)
		}
	})
}

type Repository struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Stars    int    `json:"stargazers_count"`
	Forks    int    `json:"forks_count"`
	Watchers int    `json:"watchers_count"`
	Fork     bool   `json:"fork"`
}
