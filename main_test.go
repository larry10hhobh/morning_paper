package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
	"text/tabwriter"
	"time"
)

func TestGetWeather(t *testing.T) {
	// 设置一个虚构的API密钥以进行测试
	// config.SeniverseAPIKey = "fake-api-key"

	// 调用 getWeather 函数
	daily, err := getWeather(-1, 8) // 从昨天开始，获取8天的天气预报，免费api_key还是返回3天的数据
	if err != nil {
		t.Errorf("getWeather() returned an error: %v", err)
	}
	/*	for _, day := range daily {
		// style 1
		//fmt.Printf("Date: %s, High: %s, Low: %s, Test Day: %s, Text Night: %s, precip: %s, rainfall: %s\n", day.Date, day.High, day.Low, day.TextDay, day.TextNight, day.Precip, day.Rainfall)

		// style 2
		precip, err := strconv.ParseFloat(day.Precip, 64)
		if err != nil {
			t.Errorf("getWeather() returned an error: %v", err)
			return
		}
		rainfall, err := strconv.ParseFloat(day.Rainfall, 64)
		if err != nil {
			t.Errorf("getWeather() returned an error: %v", err)
			return
		}
		date, err := time.Parse("2006-01-02", day.Date)
		if err != nil {
			t.Errorf("Error parsing date: %s", err)
			return
		}
		fmt.Println(fmt.Sprintf("%s (%s)：最高%s度，最低%s度，白天%s，晚上%s，降水概率%.2f（仅国外城市?），降水量%.2f", day.Date, date.Weekday(), day.High, day.Low, day.TextDay, day.TextNight, precip, rainfall))
	}*/

	/*	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "日期\t星期\t最高温\t最低温\t白天\t晚上\t降水概率\t降水量")

		for _, day := range daily {
			precip, err := strconv.ParseFloat(day.Precip, 64)
			if err != nil {
				t.Errorf("getWeather() returned an error: %v", err)
				return
			}
			rainfall, err := strconv.ParseFloat(day.Rainfall, 64)
			if err != nil {
				t.Errorf("getWeather() returned an error: %v", err)
				return
			}
			date, err := time.Parse("2006-01-02", day.Date)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return
			}
			weekday := date.Weekday()
			fmt.Fprintf(w, "%s\t%s\t%s度\t%s度\t%s\t%s\t%.2f\t%.2f\n", day.Date, weekday, day.High, day.Low, day.TextDay, day.TextNight, precip, rainfall)
		}

		w.Flush()*/

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "日期\t星期\t最高温度\t最低温度\t白天\t晚上\t降水概率\t降水量")

	for _, day := range daily {
		precip, err := strconv.ParseFloat(day.Precip, 32)
		if err != nil {
			t.Errorf("getWeather() returned an error: %v", err)
			return
		}
		rainfall, err := strconv.ParseFloat(day.Rainfall, 32)
		if err != nil {
			t.Errorf("getWeather() returned an error: %v", err)
			return
		}
		date, err := time.Parse("2006-01-02", day.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		weekday := date.Weekday()
		fmt.Fprintf(w, "%s\t%s\t%s度\t%s度\t%s\t%s\t%.2f\t%.2fmm\n", day.Date, weekday, day.High, day.Low, day.TextDay, day.TextNight, precip, rainfall)
	}

	w.Flush()

	tableString := buf.String()
	fmt.Println(tableString)
}
