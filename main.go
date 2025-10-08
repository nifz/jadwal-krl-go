package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nifz/jadwal-krl-go/controllers"
	"github.com/nifz/jadwal-krl-go/utils"
)

func getJadwalHandler(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("from")
	end := r.URL.Query().Get("to")

	if start == "" || end == "" {
		http.Error(w, "Parameter 'from' dan 'to' wajib diisi", http.StatusBadRequest)
		return
	}

	startStation, err := controllers.FindStationByName(start)
	if err != nil || startStation.StaID == "" {
		http.Error(w, "Stasiun keberangkatan tidak ditemukan", http.StatusNotFound)
		return
	}
	endStation, err := controllers.FindStationByName(end)
	if err != nil || endStation.StaID == "" {
		http.Error(w, "Stasiun tujuan tidak ditemukan", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Keberangkatan: %s (%s)\n", startStation.StaName, startStation.StaID)
	fmt.Fprintf(w, "Tujuan: %s (%s)\n", endStation.StaName, endStation.StaID)

	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(jakarta)
	timeFrom := now.Format("15:04")
	timeTo := now.Add(2 * time.Hour).Format("15:04")

	fmt.Fprintf(w, "Mengambil jadwal dari %s sampai %s\n", timeFrom, timeTo)
	fmt.Fprintln(w, "----------------------------------------")

	schedules, err := controllers.GetSchedule(startStation.StaID, timeFrom, timeTo)
	if err != nil {
		http.Error(w, "Gagal ambil jadwal", http.StatusInternalServerError)
		return
	}

	count := 0
	for _, sch := range schedules {
		trainSchedule, err := controllers.GetScheduleTrain(sch.TrainID)
		if err != nil {
			continue
		}

		hasStart, hasEnd := false, false
		var startTime, endTime string

		for _, stop := range trainSchedule {
			if stop.StationID == startStation.StaID {
				startTime = stop.TimeEst
				hasStart = true
			}
			if stop.StationID == endStation.StaID && hasStart {
				endTime = stop.TimeEst
				hasEnd = true
				break
			}
		}
		if hasStart && hasEnd {
			nowStr := now.Format("15:04:05")
			if (startTime < nowStr) {
				continue
			}
			waitEst, _ := utils.DurationUntil(nowStr, startTime)
			fmt.Fprintf(w, "Estimasi waktu tunggu: %s\n", waitEst)
			est, _ := utils.DurationString(startTime, endTime)
			fmt.Fprintf(w, "Kereta %s | %s\n", sch.TrainID, sch.KaName)
			fmt.Fprintf(w, "%s %s → %s %s \nEstimasi sampai: %s\n",
				startStation.StaName, startTime,
				endStation.StaName, endTime,
				est,
			)
			fmt.Fprintln(w, "----------------------------------------")

			count++
			if count == 2 {
				break
			}
		}
	}

	if count == 0 {
		fmt.Fprintln(w, "Tidak ada kereta langsung untuk rute ini dalam 2 jam ke depan.")
	}
}

func main() {
	http.HandleFunc("/jadwal-krl", getJadwalHandler)

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
