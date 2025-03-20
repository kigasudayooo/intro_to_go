package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
	UnixTime    int64  `json:"unix_time"`
	TimeZone    string `json:"timezone"`
	TimeZoneGMT string `json:"timezone_gmt"` // GMTとの時差も追加
}

func main() {
	// 日本のタイムゾーンを設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		// 日本時間で現在時刻を取得
		now := time.Now().In(jst)

		// タイムゾーンのオフセットを計算
		_, offset := now.Zone()
		gmtOffset := float64(offset) / 3600 // 秒から時間に変換

		response := TimeResponse{
			CurrentTime: now.Format(time.RFC3339),
			UnixTime:    now.Unix(),
			TimeZone:    "Asia/Tokyo",
			TimeZoneGMT: fmt.Sprintf("GMT+%.1f", gmtOffset),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
