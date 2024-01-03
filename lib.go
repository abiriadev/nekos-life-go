package nekoslifego

import (
	"encoding/json"
	"net/http"
)

const endPoint = "https://nekos.life/api/v2"

const pathTickle = "/img/tickle"
const pathSlap = "/img/slap"
const pathPoke = "/img/poke"
const pathPat = "/img/pat"
const pathNeko = "/img/neko"
const pathMeow = "/img/meow"
const pathLizard = "/img/lizard"
const pathKiss = "/img/kiss"
const pathHug = "/img/hug"
const pathFoxGirl = "/img/fox_girl"
const pathFeed = "/img/feed"
const pathCuddle = "/img/cuddle"
const pathWhy = "/why"
const pathCatText = "/cat"
const pathOwOify = "/owoify"
const pathEightBall = "/8ball"
const pathFact = "/fact"
const pathNekoGif = "/img/ngif"
const pathKemonomimi = "/img/kemonomimi"
const pathHolo = "/img/holo"
const pathSmug = "/img/smug"
const pathBaka = "/img/baka"
const pathWoof = "/img/woof"
const pathSpoiler = "/spoiler"
const pathWallpaper = "/img/wallpaper"
const pathGoose = "/img/goose"
const pathGecg = "/img/gecg"
const pathAvatar = "/img/avatar"
const pathWaifu = "/img/waifu"

type res interface {
	primary() string
}

type resUrl struct {
	url string
}

func (r resUrl) primary() string {
	return r.url
}

type resCat struct {
	cat string
}

func (r resCat) primary() string {
	return r.cat
}

type resWhy struct {
	why string
}

func (r resWhy) primary() string {
	return r.why
}

type resOwOify struct {
	owo string
}

func (r resOwOify) primary() string {
	return r.owo
}

type resFact struct {
	fact string
}

func (r resFact) primary() string {
	return r.fact
}

type EightBallResult struct {
	response string
	url string
}

func getInternal[T res](path string) (string, error) {
	res, err := http.Get(endPoint + path)
	if err != nil {
		return "", err
	}

	var body T
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.primary(), nil
}

func Tickle() (string, error) {
	return getInternal[resUrl](pathTickle)
}

func Slap() (string, error) {
	return getInternal[resUrl](pathSlap)
}

func Poke() (string, error) {
	return getInternal[resUrl](pathPoke)
}

func Pat() (string, error) {
	return getInternal[resUrl](pathPat)
}

func Neko() (string, error) {
	return getInternal[resUrl](pathNeko)
}

func Meow() (string, error) {
	return getInternal[resUrl](pathMeow)
}

func Lizard() (string, error) {
	return getInternal[resUrl](pathLizard)
}

func Kiss() (string, error) {
	return getInternal[resUrl](pathKiss)
}

func Hug() (string, error) {
	return getInternal[resUrl](pathHug)
}

func FoxGirl() (string, error) {
	return getInternal[resUrl](pathFoxGirl)
}

func Feed() (string, error) {
	return getInternal[resUrl](pathFeed)
}

func Cuddle() (string, error) {
	return getInternal[resUrl](pathCuddle)
}

func Why() (string, error) {
	return getInternal[resWhy](pathWhy)
}

func CatText() (string, error) {
	return getInternal[resCat](pathCatText)
}

func OwOify() (string, error) {
	return getInternal[resOwOify](pathOwOify)
}

func EightBall() (EightBallResult, error) {
	res, err := http.Get(endPoint + pathEightBall)
	if err != nil {
		return "", err
	}

	var body EightBallResult
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body, nil
}

func Fact() (string, error) {
	return getInternal[resFact](pathFact)
}

func NekoGif() (string, error) {
	return getInternal[resUrl](pathNekoGif)
}

func Kemonomimi() (string, error) {
	return getInternal[resUrl](pathKemonomimi)
}

func Holo() (string, error) {
	return getInternal[resUrl](pathHolo)
}

func Smug() (string, error) {
	return getInternal[resUrl](pathSmug)
}

func Baka() (string, error) {
	return getInternal[resUrl](pathBaka)
}

func Woof() (string, error) {
	return getInternal[resUrl](pathWoof)
}

func Spoiler() (string, error) {
	return getInternal[resUrl](pathSpoiler)
}

func Wallpaper() (string, error) {
	return getInternal[resUrl](pathWallpaper)
}URLSearchParams

func Goose() (string, error) {
	return getInternal[resUrl](pathGoose)
}

func Gecg() (string, error) {
	return getInternal[resUrl](pathGecg)
}

func Avatar() (string, error) {
	return getInternal[resUrl](pathAvatar)
}

func Waifu() (string, error) {
	return getInternal[resUrl](pathWaifu)
}
