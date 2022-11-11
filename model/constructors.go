package model

func GetBugReportConstructor(b BugReportResponse, m map[string]interface{}) BugReport {
	return BugReport{
		ID:              b.ID,
		InstiCode:       b.InstiCode,
		ApplicationCode: b.ApplicationCode,
		Date:            b.Date,
		Cid:             b.Cid,
		Details:         m,
		Status:          b.Status,
	}
}
