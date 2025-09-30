package utils

import (
	"fmt"
	"time"
)

func DurationString(start, end string) (string, error) {
	layout := "15:04:05"
	t1, err := time.Parse(layout, start)
	if err != nil {
		return "", err
	}
	t2, err := time.Parse(layout, end)
	if err != nil {
		return "", err
	}

	if t2.Before(t1) {
		// kalau jam tujuan lebih kecil, berarti sudah lewat tengah malam
		t2 = t2.Add(24 * time.Hour)
	}

	diff := t2.Sub(t1)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) % 60

	if hours > 0 {
		return fmt.Sprintf("%d jam %d menit", hours, minutes), nil
	}
	return fmt.Sprintf("%d menit", minutes), nil
}

func DurationUntil(now, target string) (string, error) {
	layout := "15:04:05"
	tNow, err := time.Parse(layout, now)
	if err != nil {
		return "", err
	}
	tTarget, err := time.Parse(layout, target)
	if err != nil {
		return "", err
	}

	if tTarget.Before(tNow) {
		tTarget = tTarget.Add(24 * time.Hour)
	}

	diff := tTarget.Sub(tNow)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) % 60

	if hours > 0 {
		return fmt.Sprintf("%d jam %d menit", hours, minutes), nil
	}
	return fmt.Sprintf("%d menit", minutes), nil
}
