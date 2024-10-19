package presentation

import (
	"fmt"
	"go-stress-tester/internal/usecase"
	"strings"
	"time"
)

type ReportBuilder struct{}

func NewReportBuilder() *ReportBuilder {
	return &ReportBuilder{}
}

func (b *ReportBuilder) BuildPresentableReport(result usecase.ResultDto) string {
	var sb strings.Builder

	duration := time.Duration(result.TotalExecutionTimeInMilliseconds) * time.Millisecond
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	milliseconds := int(duration.Milliseconds()) % 1000

	sb.WriteString("\nStress Test Report\n")
	sb.WriteString("\n=======================================================\n")
	sb.WriteString(fmt.Sprintf("Total Execution Time: %02dm %02ds %03dms\n", minutes, seconds, milliseconds))
	sb.WriteString(fmt.Sprintf("Total Requests: %d\n", result.TotalRequests))
	sb.WriteString("Successful Requests:\n")
	for code, count := range result.SuccessFullRequests {
		sb.WriteString(fmt.Sprintf("  %d: %d\n", code, count))
	}
	sb.WriteString("Error Requests:\n")
	for code, count := range result.ErrorRequests {
		sb.WriteString(fmt.Sprintf("  %d: %d\n", code, count))
	}

	return sb.String()
}
