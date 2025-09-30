package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/nifz/jadwal-krl-go/dtos"
	"github.com/nifz/jadwal-krl-go/utils"
)

func FindStationByName(name string) (dtos.Station, error) {
	// Normalisasi input
	nameNorm := strings.ToUpper(name)
	if nameNorm == "" {
		return dtos.Station{}, errors.New("input station name is empty")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-partner.krl.co.id/krl-webs/v1/krl-station", nil)
	if err != nil {
		return dtos.Station{}, err
	}

	token, err := utils.Token()
	if err != nil {
		return dtos.Station{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Host", "api-partner.krl.co.id")
	req.Header.Set("Origin", "https://commuterline.id")
	req.Header.Set("Referer", "https://commuterline.id/")

	resp, err := client.Do(req)
	if err != nil {
		return dtos.Station{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	log.Println("DEBUG STATUS:", resp.StatusCode)
	log.Println("DEBUG BODY:", string(body))
	if err != nil {
		return dtos.Station{}, err
	}

	var result dtos.KrlResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return dtos.Station{}, err
	}

	// --- Exact match only ---
	for _, s := range result.Data {
		staNorm := strings.ToUpper(s.StaName)
		if staNorm == nameNorm {
			return s, nil
		}
	}

	return dtos.Station{}, errors.New("station not found")
}
