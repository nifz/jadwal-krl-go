package controllers

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/nifz/jadwal-krl-go/dtos"
	"github.com/nifz/jadwal-krl-go/utils"
)

func FindStationByName(name string) (dtos.Station, error) {
	nameNorm := strings.ToUpper(strings.TrimSpace(name))
	if nameNorm == "" {
		return dtos.Station{}, errors.New("input station name is empty")
	}

	token, err := utils.Token()
	if err != nil {
		return dtos.Station{}, err
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("User-Agent", "Mozilla/5.0").
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Origin", "https://commuterline.id").
		SetHeader("Referer", "https://commuterline.id/").
		Get("https://api-partner.krl.co.id/krl-webs/v1/krl-station")

	if err != nil {
		return dtos.Station{}, err
	}

	if resp.StatusCode() != 200 {
		return dtos.Station{}, errors.New("API error: " + resp.Status())
	}

	var result dtos.KrlResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dtos.Station{}, err
	}

	for _, s := range result.Data {
		if strings.ToUpper(s.StaName) == nameNorm {
			return s, nil
		}
	}

	return dtos.Station{}, errors.New("station not found")
}
