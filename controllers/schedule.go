package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/nifz/jadwal-krl-go/dtos"
	"github.com/nifz/jadwal-krl-go/utils"
)

// Function untuk ambil jadwal KRL berdasarkan station ID & range waktu
func GetSchedule(stationID, timeFrom, timeTo string) ([]dtos.Schedule, error) {
	if stationID == "" {
		return nil, errors.New("stationID is required")
	}
	if timeFrom == "" || timeTo == "" {
		return nil, errors.New("timeFrom and timeTo are required")
	}

	url := fmt.Sprintf(
		"https://api-partner.krl.co.id/krl-webs/v1/schedule?stationid=%s&timefrom=%s&timeto=%s",
		stationID, timeFrom, timeTo,
	)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token, err := utils.Token()
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Host", "api-partner.krl.co.id")
	req.Header.Set("Origin", "https://commuterline.id")
	req.Header.Set("Referer", "https://commuterline.id/")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Debug kalau perlu
	// fmt.Println("RAW RESPONSE:", string(body))

	var result dtos.ScheduleResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.Status != 200 {
		return nil, errors.New("failed to fetch schedule: " + fmt.Sprint(result.Status))
	}

	return result.Data, nil
}
