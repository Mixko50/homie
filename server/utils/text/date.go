package text

import "time"

func FormatTime(t time.Time) time.Time {
	timeString := t.String()
	parsed, _ := time.Parse("2006-01-02 15:04:05", timeString)
	return parsed
}

func FormatTimeToString(t *time.Time) *string {
	parsed := t.Format("2006-01-02 15:04:05")
	return &parsed
}
