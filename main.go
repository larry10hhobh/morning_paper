package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"text/tabwriter"
	"time"
)

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config.yaml file: %v", err)
	}
}

func main() {

	// Get Telegram bot
	bot, err := getBotWithSocks5Proxy()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bot username: %s\n", bot.Self.UserName)
	fmt.Printf("Bot ID: %d\n", bot.Self.ID)

	// Fetch weather and news
	daily, err := getWeather(-1, 8)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "日期\t星期\t最高温度\t最低温度\t白天\t晚上\t降水概率\t降水量")

	for _, day := range daily {
		precip, err := strconv.ParseFloat(day.Precip, 32)
		if err != nil {
			return
		}
		rainfall, err := strconv.ParseFloat(day.Rainfall, 32)
		if err != nil {
			return
		}
		date, err := time.Parse("2006-01-02", day.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		weekday := date.Weekday()
		fmt.Fprintf(w, "%s\t%s\t%s度\t%s度\t%s\t%s\t%.1f\t%.1fmm\n", day.Date, weekday, day.High, day.Low, day.TextDay, day.TextNight, precip, rainfall)
	}
	w.Flush()
	tableString := buf.String()
	fmt.Println(tableString)

	// news := getNews()

	// Send weather and news to Telegram
	sendToTelegram(tableString, bot)
}
