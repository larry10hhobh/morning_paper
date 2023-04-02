package main

import (
	"fmt"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// 设置一个虚构的API密钥以进行测试
	// config.SeniverseAPIKey = "fake-api-key"

	// 调用 getWeather 函数
	daily, err := getWeather(-1, 8) // 从昨天开始，获取8天的天气预报，免费api_key还是返回3天的数据
	if err != nil {
		t.Errorf("getWeather() returned an error: %v", err)
	}
	for _, day := range daily {
		fmt.Printf("Date: %s, High: %s, Low: %s, Test Day: %s, Text Night: %s, precip: %s, rainfall: %s\n", day.Date, day.High, day.Low, day.TextDay, day.TextNight, day.Precip, day.Rainfall)
	}
}
