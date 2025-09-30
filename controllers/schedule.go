package controllers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
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

	token, err := utils.Token()
	if err != nil {
		return nil, err
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("User-Agent", "Mozilla/5.0").
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Origin", "https://commuterline.id").
		SetHeader("Referer", "https://commuterline.id/").
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API error: %s", resp.Status())
	}

	var result dtos.ScheduleResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	if result.Status != 200 {
		return nil, errors.New("failed to fetch schedule: " + fmt.Sprint(result.Status))
	}

	return result.Data, nil
}
