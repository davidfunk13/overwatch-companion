// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type MutateItemPayload interface {
	IsMutateItemPayload()
}

type OptionalDataPayload interface {
	IsOptionalDataPayload()
}

type Battletag struct {
	ID          int      `json:"id"`
	UserID      string   `json:"userId"`
	Name        string   `json:"name"`
	URLName     string   `json:"urlName"`
	BlizzID     int      `json:"blizzId"`
	Level       int      `json:"level"`
	PlayerLevel int      `json:"playerLevel"`
	Platform    Platform `json:"platform"`
	IsPublic    *bool    `json:"isPublic"`
	Portrait    string   `json:"portrait"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   *string  `json:"updated_at"`
}

func (Battletag) IsOptionalDataPayload() {}

type BlizzBattletag struct {
	Name        string   `json:"name"`
	URLName     string   `json:"urlName"`
	BlizzID     int      `json:"blizzId"`
	Level       int      `json:"level"`
	PlayerLevel int      `json:"playerLevel"`
	Platform    Platform `json:"platform"`
	IsPublic    bool     `json:"isPublic"`
	Portrait    string   `json:"portrait"`
}

type Game struct {
	ID           int          `json:"id"`
	UserID       string       `json:"userId"`
	BattletagID  int          `json:"battletagId"`
	SessionID    int          `json:"sessionId"`
	Location     Location     `json:"location"`
	Role         Role         `json:"role"`
	SrIn         int          `json:"sr_in"`
	SrOut        int          `json:"sr_out"`
	MatchOutcome MatchOutcome `json:"match_outcome"`
	CreatedAt    string       `json:"created_at"`
	UpdatedAt    *string      `json:"updated_at"`
}

func (Game) IsOptionalDataPayload() {}

type InputBattletag struct {
	UserID      string   `json:"userId"`
	Name        string   `json:"name"`
	URLName     string   `json:"urlName"`
	BlizzID     int      `json:"blizzId"`
	Level       int      `json:"level"`
	PlayerLevel int      `json:"playerLevel"`
	Platform    Platform `json:"platform"`
	IsPublic    bool     `json:"isPublic"`
	Portrait    string   `json:"portrait"`
}

type InputGame struct {
	UserID       string       `json:"userId"`
	BattletagID  int          `json:"battletagId"`
	SessionID    int          `json:"sessionId"`
	Location     Location     `json:"location"`
	Role         Role         `json:"role"`
	SrOut        int          `json:"sr_out"`
	MatchOutcome MatchOutcome `json:"match_outcome"`
}

type InputGetGame struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	BattletagID int    `json:"battletagId"`
	SessionID   int    `json:"sessionId"`
	Role        *Role  `json:"role"`
}

type InputGetGames struct {
	UserID      string `json:"userId"`
	BattletagID int    `json:"battletagId"`
	SessionID   int    `json:"sessionId"`
	Role        *Role  `json:"role"`
}

type InputGetMostRecentSession struct {
	UserID      string `json:"userId"`
	BattletagID int    `json:"battletagId"`
}

type InputGetOneBattletag struct {
	UserID      string `json:"userId"`
	BattletagID int    `json:"battletagId"`
}

type InputGetOneSession struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	BattletagID int    `json:"battletagId"`
}

type InputGetSessions struct {
	UserID      string `json:"userId"`
	BattletagID int    `json:"battletagId"`
}

type InputSession struct {
	UserID            string `json:"userId"`
	BattletagID       int    `json:"battletagId"`
	StartingSrTank    *int   `json:"starting_sr_tank"`
	StartingSrDamage  *int   `json:"starting_sr_damage"`
	StartingSrSupport *int   `json:"starting_sr_support"`
}

type InputUpdateSessionStartingSr struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	BattletagID int    `json:"battletagId"`
	Role        Role   `json:"role"`
	StartingSr  int    `json:"starting_sr"`
}

type MutateItemPayloadFailure struct {
	ID      int     `json:"id"`
	Success bool    `json:"success"`
	Error   string  `json:"error"`
	Data    *string `json:"data"`
}

func (MutateItemPayloadFailure) IsMutateItemPayload() {}

type MutateItemPayloadSuccess struct {
	ID      int                 `json:"id"`
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    OptionalDataPayload `json:"data"`
}

func (MutateItemPayloadSuccess) IsMutateItemPayload() {}

type Session struct {
	ID                int     `json:"id"`
	UserID            string  `json:"userId"`
	BattletagID       int     `json:"battletagId"`
	StartingSrTank    int     `json:"starting_sr_tank"`
	SrTank            int     `json:"sr_tank"`
	StartingSrDamage  int     `json:"starting_sr_damage"`
	SrDamage          int     `json:"sr_damage"`
	StartingSrSupport int     `json:"starting_sr_support"`
	SrSupport         int     `json:"sr_support"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

func (Session) IsOptionalDataPayload() {}

type Location string

const (
	LocationBusan               Location = "BUSAN"
	LocationIlios               Location = "ILIOS"
	LocationLijiangtower        Location = "LIJIANGTOWER"
	LocationNepal               Location = "NEPAL"
	LocationOasis               Location = "OASIS"
	LocationHanamura            Location = "HANAMURA"
	LocationTempleofanubis      Location = "TEMPLEOFANUBIS"
	LocationVolskayaindustries  Location = "VOLSKAYAINDUSTRIES"
	LocationDorado              Location = "DORADO"
	LocationHavana              Location = "HAVANA"
	LocationJunkertown          Location = "JUNKERTOWN"
	LocationRialto              Location = "RIALTO"
	LocationRoute66             Location = "ROUTE66"
	LocationWatchpointgibraltar Location = "WATCHPOINTGIBRALTAR"
	LocationBlizzardworld       Location = "BLIZZARDWORLD"
	LocationEichenwalde         Location = "EICHENWALDE"
	LocationHollywood           Location = "HOLLYWOOD"
	LocationKingsrow            Location = "KINGSROW"
	LocationNumbani             Location = "NUMBANI"
)

var AllLocation = []Location{
	LocationBusan,
	LocationIlios,
	LocationLijiangtower,
	LocationNepal,
	LocationOasis,
	LocationHanamura,
	LocationTempleofanubis,
	LocationVolskayaindustries,
	LocationDorado,
	LocationHavana,
	LocationJunkertown,
	LocationRialto,
	LocationRoute66,
	LocationWatchpointgibraltar,
	LocationBlizzardworld,
	LocationEichenwalde,
	LocationHollywood,
	LocationKingsrow,
	LocationNumbani,
}

func (e Location) IsValid() bool {
	switch e {
	case LocationBusan, LocationIlios, LocationLijiangtower, LocationNepal, LocationOasis, LocationHanamura, LocationTempleofanubis, LocationVolskayaindustries, LocationDorado, LocationHavana, LocationJunkertown, LocationRialto, LocationRoute66, LocationWatchpointgibraltar, LocationBlizzardworld, LocationEichenwalde, LocationHollywood, LocationKingsrow, LocationNumbani:
		return true
	}
	return false
}

func (e Location) String() string {
	return string(e)
}

func (e *Location) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Location(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Location", str)
	}
	return nil
}

func (e Location) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MatchOutcome string

const (
	MatchOutcomeLoss MatchOutcome = "LOSS"
	MatchOutcomeWin  MatchOutcome = "WIN"
	MatchOutcomeDraw MatchOutcome = "DRAW"
)

var AllMatchOutcome = []MatchOutcome{
	MatchOutcomeLoss,
	MatchOutcomeWin,
	MatchOutcomeDraw,
}

func (e MatchOutcome) IsValid() bool {
	switch e {
	case MatchOutcomeLoss, MatchOutcomeWin, MatchOutcomeDraw:
		return true
	}
	return false
}

func (e MatchOutcome) String() string {
	return string(e)
}

func (e *MatchOutcome) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MatchOutcome(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MatchOutcome", str)
	}
	return nil
}

func (e MatchOutcome) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Platform string

const (
	PlatformPc             Platform = "PC"
	PlatformNintendoswitch Platform = "NINTENDOSWITCH"
	PlatformXbox           Platform = "XBOX"
	PlatformPlaystation    Platform = "PLAYSTATION"
)

var AllPlatform = []Platform{
	PlatformPc,
	PlatformNintendoswitch,
	PlatformXbox,
	PlatformPlaystation,
}

func (e Platform) IsValid() bool {
	switch e {
	case PlatformPc, PlatformNintendoswitch, PlatformXbox, PlatformPlaystation:
		return true
	}
	return false
}

func (e Platform) String() string {
	return string(e)
}

func (e *Platform) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Platform(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Platform", str)
	}
	return nil
}

func (e Platform) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleTank    Role = "TANK"
	RoleDamage  Role = "DAMAGE"
	RoleSupport Role = "SUPPORT"
)

var AllRole = []Role{
	RoleTank,
	RoleDamage,
	RoleSupport,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleTank, RoleDamage, RoleSupport:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
