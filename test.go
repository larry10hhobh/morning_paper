package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "strings"
)

type Weather struct {
    Temperature string `json:"temperature"`
    Conditions  string `json:"conditions"`
}

type News struct {
    Title string `json:"title"`
    Url   string `json:"url"`
}

func main() {
    weather := getWeather()
    news := getNews()

    // Replace with your own WeChat or SMS API
    sendToWeChat("YOUR_WECHAT_ID", weather.Temperature, news.Title)
    sendToSms("YOUR_PHONE_NUMBER", weather.Temperature, news.Title)
}

func getWeather() Weather {
    // Replace with your own weather API
    resp, err := http.Get("https://api.example.com/weather")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    var weather Weather
    err = json.Unmarshal(body, &weather)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    return weather
}

func getNews() News {
    // Replace with your own news API
    resp, err := http.Get("https://api.example.com/news")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    var news []News
    err = json.Unmarshal(body, &news)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    return news[0]
}

func sendToWeChat(wechatId string, weather string, news string) {
    // Replace with your own WeChat API
    api := "https://api.example.com/wechat/send"

    data := url.Values{}
    data.Set("wechatId", wechatId)
    data.Set("message", fmt.Sprintf("Temperature: %s\nNews: %s", weather, news))

    req, err := http.NewRequest("POST", api, strings.NewReader(data.Encode()))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(body))
}

func sendToSms(phoneNumber string, weather string, news string) {
    // Replace with your own SMS API
    api := "https://api.example.com/sms/send"

    data := url.Values{}
    data.Set("phoneNumber", phoneNumber)
    data.Set("message", fmt.Sprintf("Temperature: %s\nNews: %s", weather, news))

    req, err := http.NewRequest("POST", api, strings.NewReader(data.Encode()))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(body)
}

