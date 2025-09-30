package controllers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/nifz/jadwal-krl-go/dtos"
	"github.com/nifz/jadwal-krl-go/utils"
)

// Ambil detail perjalanan train berdasarkan trainID
func GetScheduleTrain(trainID string) ([]dtos.ScheduleTrain, error) {
	if trainID == "" {
		return nil, errors.New("trainID is required")
	}

	url := fmt.Sprintf("https://api-partner.krl.co.id/krl-webs/v1/schedule-train?trainid=%s", trainID)

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

	var result dtos.ScheduleTrainResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	if result.Status != 200 {
		return nil, errors.New("failed to fetch train schedule detail")
	}

	return result.Data, nil
}
