package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/nifz/jadwal-krl-go/dtos"
)

func FindStationByName(name string) (dtos.Station, error) {
	// Normalisasi input
	nameNorm := strings.ToUpper(name)
	if nameNorm == "" {
		return dtos.Station{}, errors.New("input station name is empty")
	}

	// Fetch API eksternal
	resp, err := http.Get("https://api-partner.krl.co.id/krl-webs/v1/krl-station")
	if err != nil {
		return dtos.Station{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
