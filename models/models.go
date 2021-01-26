package models

// User : Model struct for User.
type User struct {
	ID         int    `json:"id"`
	Battletag  string `json:"battletag"`
	Identifier int    `json:"Identifier"`
	Email      string `json:"email"`
}

// Hero : Model struct for hero.
type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
}
