package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

func SearchBattletags(b string) []model.BlizzBattletag {

	url := "https://playoverwatch.com/en-us/search/account-by-name/" + b
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	defer res.Body.Close()

	// We can probably do this WAY better.....
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	var data []map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	var battletags []model.BlizzBattletag

	for _, b := range data {

		p := b["platform"].(string)
		pSanitized := strings.Join(strings.Split(p, " "), "")
		pUpper := strings.ToUpper(pSanitized)

		t := model.BlizzBattletag{
			Name:        b["name"].(string),
			URLName:     b["urlName"].(string),
			BlizzID:     int(b["id"].(float64)),
			Level:       int(b["level"].(float64)),
			PlayerLevel: int(b["playerLevel"].(float64)),
			Platform:    model.Platform(pUpper),
			IsPublic:    b["isPublic"].(bool),
			Portrait:    b["portrait"].(string),
		}

		battletags = append(battletags, t)
	}

	return battletags
}
