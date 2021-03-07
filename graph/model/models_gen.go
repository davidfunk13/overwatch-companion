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

type Battletag struct {
	ID          int      `json:"id"`
	UserID      int      `json:"userId"`
	Name        string   `json:"name"`
	URLName     string   `json:"urlName"`
	BlizzID     int      `json:"blizzId"`
	Level       int      `json:"level"`
	PlayerLevel int      `json:"playerLevel"`
	Platform    Platform `json:"platform"`
	IsPublic    *bool    `json:"isPublic"`
	Portrait    string   `json:"portrait"`
}

type Game struct {
	ID       int `json:"id"`
	UserID   int `json:"userId"`
	SessonID int `json:"sessonId"`
}

type InputBattletag struct {
	UserID      int      `json:"userId"`
	Name        string   `json:"name"`
	URLName     string   `json:"urlName"`
	BlizzID     int      `json:"blizzId"`
	Level       int      `json:"level"`
	PlayerLevel int      `json:"playerLevel"`
	Platform    Platform `json:"platform"`
	IsPublic    *bool    `json:"isPublic"`
	Portrait    string   `json:"portrait"`
}

type InputGame struct {
	UserID    int `json:"userId"`
	SessionID int `json:"sessionId"`
}

type InputSession struct {
	UserID   int  `json:"userId"`
	RoleType Role `json:"roleType"`
}

type MutateItemPayloadFailure struct {
	ID      int    `json:"id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func (MutateItemPayloadFailure) IsMutateItemPayload() {}

type MutateItemPayloadSuccess struct {
	ID      int    `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (MutateItemPayloadSuccess) IsMutateItemPayload() {}

type Session struct {
	ID       int  `json:"id"`
	UserID   int  `json:"userId"`
	RoleType Role `json:"roleType"`
}

type Platform string

const (
	PlatformPc          Platform = "PC"
	PlatformNintendo    Platform = "NINTENDO"
	PlatformXbox        Platform = "XBOX"
	PlatformPlaystation Platform = "PLAYSTATION"
)

var AllPlatform = []Platform{
	PlatformPc,
	PlatformNintendo,
	PlatformXbox,
	PlatformPlaystation,
}

func (e Platform) IsValid() bool {
	switch e {
	case PlatformPc, PlatformNintendo, PlatformXbox, PlatformPlaystation:
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
