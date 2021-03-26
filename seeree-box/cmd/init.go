/*
Copyright © 2021 Haseb Ansari ansari-haseb GitHub

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"seeree-box/tv"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var TMDB_API_KEY_v3 string = "api-key"
var TMDB_API_KEY_v4 string = "bearer token" // Used as header "Authorizazion Bearer <<access_token>>"
var TMDB_API_BASE_URL string = "https://api.themoviedb.org/3/"

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "The command initializes the toolkit and gets ready to take user input.",
	Long:  `seeree-box init starts a user interactive flow to view any episode's description based on TMDB APIs https://developers.themoviedb.org/3/getting-started/introduction`,
	Run: func(cmd *cobra.Command, args []string) {

		var languages []string = []string{"en-US", "de-DE"}

		selectedLanguage := promptSelector(2, languages, "Choose your language: ")

		if selectedLanguage == "en-US" {
			fmt.Println("Enter the TV Show Name: ")
		} else {
			fmt.Println("Geben Sie die Namen der TV Show ein: ")
		}

		var tvShowName string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			tvShowName = scanner.Text()
		}
		var tvShowSeries tv.TvSeries = findShows(tvShowName, "1", selectedLanguage) // Finds TV Shows with Page 1 of search result - as it is a first step
		if tvShowSeries.TotalResults == 0 {
			fmt.Println("No result found for your input. Please re-run the command with different TV Show title.....") // When there is no result for the given title then CLI comes out show the message-
			return
		}
		processNavigation(tvShowSeries, tvShowName, selectedLanguage)
	},
}

// Finds TV shows for the given TV show title and pageNo
func findShows(tvShowName string, pageNo string, locale string) tv.TvSeries {
	tvShowName = strings.ReplaceAll(tvShowName, " ", "+")
	EXTENDED_URL := "search/tv?api_key=" + TMDB_API_KEY_v3 + "&query=" + tvShowName + "&page=" + pageNo + "&language=" + locale
	var tvShowSeries string = getRequestToTMDB(EXTENDED_URL)
	var tvSeries tv.TvSeries
	json.Unmarshal([]byte(tvShowSeries), &tvSeries)
	return tvSeries
}

// Cases to process, depending on the User's select input. If pagination links {FIRST, NEXT, etc.} are selected then the result pages are changed otherwise, the TV Show details will be triggered on selecting a show
func navigatorAction(navAction string, showName string, pageNo int, lastPageNo int, showID int, locale string) {
	var payload tv.TvSeries
	switch navAction {
	case "Next":
		payload = findShows(showName, strconv.Itoa(pageNo+1), locale)
		processNavigation(payload, showName, locale)
	case "Prev":
		payload = findShows(showName, strconv.Itoa(pageNo-1), locale)
		processNavigation(payload, showName, locale)
	case "First":
		payload = findShows(showName, strconv.Itoa(1), locale)
		processNavigation(payload, showName, locale)
	case "Last":
		payload = findShows(showName, strconv.Itoa(lastPageNo), locale)
		processNavigation(payload, showName, locale)
	default:
		processDefaultCase(showID, locale)
	}
}

func processDefaultCase(showID int, locale string) {
	var seasonsDetails tv.TvShowDetails = getShowSeasons(showID, locale)
	var seasons []string
	for i := 1; i <= seasonsDetails.NumberOfSeasons; i++ {
		seasons = append(seasons, fmt.Sprintf("Season %d", i))
	}
	label := getLabel(locale, "Season")
	res := promptSelector(15, seasons, label)
	res = strings.Replace(res, "Season ", "", 1)
	prepareEpisodesListPrompt(res, strconv.Itoa(showID), locale)
}

func getLabel(locale string, selector string) string {
	if locale == "en-US" {
		return "Select your " + selector
	}
	return "Wahlen Sie Ihre " + selector
}

func processNavigation(tvShowSeries tv.TvSeries, showName string, locale string) {
	var tvshowsResults []tv.Results = tvShowSeries.Results
	var seriesSelectNames []string

	for _, s := range tvshowsResults {
		seriesSelectNames = append(seriesSelectNames, s.Name)
	}

	var navigator []string = displayNav(tvShowSeries.Page, tvShowSeries.TotalPages, tvShowSeries.TotalResults)
	seriesSelectNames = append(seriesSelectNames, navigator...)

	label := getLabel(locale, "TV Show")
	res := promptSelector(20, seriesSelectNames, label)

	var showID int
	for _, rs := range tvshowsResults {
		if rs.Name == res {
			showID = rs.ID
		}
	}
	navigatorAction(res, showName, tvShowSeries.Page, tvShowSeries.TotalPages, showID, locale) // Calls the switch case to be solved with the "navAction" text --> "res" in this case
}

// The method dynamically appends elements to the array when required for pagination. For e.g; 1st page will show {NEXT and LAST} links, middle pages will show {FIRST, NEXT, PREV, LAST} links and so on
func displayNav(pageNo int, totalPages int, totalResults int) []string {
	if pageNo == 1 && totalPages > 1 {
		return []string{"Next", "Last"}
	}
	if pageNo > 1 && pageNo < totalPages {
		return []string{"First", "Prev", "Next", "Last"}
	}
	if pageNo == totalPages && totalResults > 20 {
		return []string{"Prev", "First"}
	}
	return []string{}
}

// GET Show details using the show ID
func getShowSeasons(showID int, locale string) tv.TvShowDetails {
	EXTENDED_URL := "tv/" + strconv.Itoa(showID) + "?api_key=" + TMDB_API_KEY_v3 + "&language=" + locale
	var getShowDetails string = getRequestToTMDB(EXTENDED_URL)
	var tvShowDetails tv.TvShowDetails
	json.Unmarshal([]byte(getShowDetails), &tvShowDetails)
	return tvShowDetails
}

// Common GET request to TMDB APIs with customized EXTENDED_URL
func getRequestToTMDB(EXTENDED_URL string) string {
	resp, err := http.Get(TMDB_API_BASE_URL + EXTENDED_URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb
}

// GET Season Details with the season number and show ID
func getSeasonDetails(season string, showID string, locale string) tv.TvSeriesSeasonDetails {
	EXTENDED_URL := "tv/" + showID + "/season/" + season + "?api_key=" + TMDB_API_KEY_v3 + "&language=" + locale
	var getSeasonDetails string = getRequestToTMDB(EXTENDED_URL)
	var tvSeasonDetails tv.TvSeriesSeasonDetails
	json.Unmarshal([]byte(getSeasonDetails), &tvSeasonDetails)
	return tvSeasonDetails
}

// Season details and preparing list of Episodes in that Season
func prepareEpisodesListPrompt(seasonNo string, showID string, locale string) {
	var seasonDetails tv.TvSeriesSeasonDetails = getSeasonDetails(seasonNo, showID, locale)
	var episodes []string
	var counter int = 0
	for _, episode := range seasonDetails.Episodes {
		counter = counter + 1
		episodes = append(episodes, fmt.Sprintf("Episode - %d  -->  %s", counter, episode.Name))
	}
	label := getLabel(locale, "Episode")
	res := promptSelector(15, episodes, label)
	episodeName := res[18:len(res)]

	// Displaying final result of the Episode Overview
	for _, episode := range seasonDetails.Episodes {
		if strings.TrimSpace(episodeName) == episode.Name {
			fmt.Println("##########################################")
			fmt.Println("EPISODE - ", episode.EpisodeNumber)
			fmt.Println("##########################################")
			fmt.Println("Title: ", episode.Name)
			fmt.Println("Summary: ", episode.Overview)
		}
	}
}

func promptSelector(size int, elements []string, label string) string {
	prompt := promptui.Select{ // Displays an interactive select list tool
		Label: label,
		Size:  size,
		Items: elements,
		Templates: &promptui.SelectTemplates{
			Active:   ` 🎬 {{ . | cyan | bold }}`,
			Inactive: `   {{ . | cyan }}`,
			Selected: `{{ "✔" | green | bold }} {{ "You Selected: " | bold }}: {{ . | cyan }}`,
		},
	}
	_, res, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "Something failed while initializing PromptUI Select."
	}
	return res
}

func init() {
	rootCmd.AddCommand(initCmd)
}
