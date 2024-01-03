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

func (r *resUrl) primary() string {
	return r.url
}

type resCat struct {
	cat string
}

func (r *resCat) primary() string {
	return r.cat
}

type resWhy struct {
	why string
}

func (r *resWhy) primary() string {
	return r.why
}

type resOwOify struct {
	owo string
}

func (r *resOwOify) primary() string {
	return r.owo
}

type resFact struct {
	fact string
}

func (r *resFact) primary() string {
	return r.fact
}

func Tickle() (string, error) {
	res, err := http.Get(endPoint + pathTickle)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Slap() (string, error) {
	res, err := http.Get(endPoint + pathSlap)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Poke() (string, error) {
	res, err := http.Get(endPoint + pathPoke)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Pat() (string, error) {
	res, err := http.Get(endPoint + pathPat)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Neko() (string, error) {
	res, err := http.Get(endPoint + pathNeko)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Meow() (string, error) {
	res, err := http.Get(endPoint + pathMeow)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Lizard() (string, error) {
	res, err := http.Get(endPoint + pathLizard)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Kiss() (string, error) {
	res, err := http.Get(endPoint + pathKiss)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Hug() (string, error) {
	res, err := http.Get(endPoint + pathHug)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func FoxGirl() (string, error) {
	res, err := http.Get(endPoint + pathFoxGirl)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Feed() (string, error) {
	res, err := http.Get(endPoint + pathFeed)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Cuddle() (string, error) {
	res, err := http.Get(endPoint + pathCuddle)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Why() (string, error) {
	res, err := http.Get(endPoint + pathWhy)
	if err != nil {
		return "", err
	}

	var body resWhy
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.why, nil
}

func CatText() (string, error) {
	res, err := http.Get(endPoint + pathCatText)
	if err != nil {
		return "", err
	}

	var body resCat
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.cat, nil
}

func OwOify() (string, error) {
	res, err := http.Get(endPoint + pathOwOify)
	if err != nil {
		return "", err
	}

	var body resOwOify
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.owo, nil
}

func EightBall() (string, error) {
	res, err := http.Get(endPoint + pathEightBall)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Fact() (string, error) {
	res, err := http.Get(endPoint + pathFact)
	if err != nil {
		return "", err
	}

	var body resFact
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.fact, nil
}

func NekoGif() (string, error) {

	res, err := http.Get(endPoint + pathNekoGif)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Kemonomimi() (string, error) {

	res, err := http.Get(endPoint + pathKemonomimi)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Holo() (string, error) {
	res, err := http.Get(endPoint + pathHolo)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Smug() (string, error) {
	res, err := http.Get(endPoint + pathSmug)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Baka() (string, error) {
	res, err := http.Get(endPoint + pathBaka)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Woof() (string, error) {
	res, err := http.Get(endPoint + pathWoof)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Spoiler() (string, error) {
	res, err := http.Get(endPoint + pathSpoiler)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Wallpaper() (string, error) {
	res, err := http.Get(endPoint + pathWallpaper)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Goose() (string, error) {
	res, err := http.Get(endPoint + pathGoose)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Gecg() (string, error) {
	res, err := http.Get(endPoint + pathGecg)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Avatar() (string, error) {
	res, err := http.Get(endPoint + pathAvatar)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}

func Waifu() (string, error) {
	res, err := http.Get(endPoint + pathWaifu)
	if err != nil {
		return "", err
	}

	var url resUrl
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.url, nil
}
