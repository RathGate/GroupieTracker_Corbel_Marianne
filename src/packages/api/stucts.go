package api

type FullRequest struct {
	Data struct {
		Creatures struct {
			Food    []Item `json:"food"`
			NonFood []Item `json:"non_food"`
		} `json:"creatures"`
		Equipment []Item `json:"equipment"`
		Materials []Item `json:"materials"`
		Monsters  []Item `json:"monsters"`
		Treasure  []Item `json:"treasure"`
	} `json:"data"`
}

type CategoryRequest struct {
	Items []Item `json:"data"`
}

type CreaturesRequest struct {
	Data struct {
		Food    []Item `json:"food"`
		NonFood []Item `json:"non_food"`
	} `json:"data"`
}

type EntryRequest struct {
	Item Item `json:"data"`
}

type Item struct {
	Attack          int      `json:"attack"`
	Category        string   `json:"category"`
	CommonLocations []string `json:"common_locations"`
	CookingEffect   string   `json:"cooking_effect"`
	Defense         int      `json:"defense"`
	Food            bool
	HeartsRecovered float64  `json:"hearts_recovered"`
	Description     string   `json:"description"`
	Drops           []string `json:"drops"`
	ID              int      `json:"id"`
	Image           string   `json:"image"`
	MasterExclusive bool
	MasterID        int
	DisplayMaster   bool
	Name            string `json:"name"`
}
