package app

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lameaux/bro/internal/client/config"
	"github.com/lameaux/bro/internal/client/stats"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	latencyPercentile = 99
)

func resultsTable(conf *config.Config, results *stats.Stats, success bool) string {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("Name: %s\nPath: %s\n", conf.Name, conf.FileName))

	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(&output)
	tableWriter.AppendHeader(table.Row{
		"Scenario", "Total", "Success", "Failed", "Timeout", "Invalid", "Latency @P99", "Duration", "RPS", "Passed",
	})

	for _, scenario := range conf.Scenarios {
		counters := results.RequestCounters(scenario.Name)
		if counters == nil {
			log.Warn().
				Dict("scenario", zerolog.Dict().Str("name", scenario.Name)).
				Msg("missing stats")

			continue
		}

		tableWriter.AppendRow(table.Row{
			scenario.Name,
			counters.Counter(stats.CounterTotal),
			counters.Counter(stats.CounterSuccess),
			counters.Counter(stats.CounterFailed),
			counters.Counter(stats.CounterTimeout),
			counters.Counter(stats.CounterInvalid),
			fmt.Sprintf("%d ms", counters.LatencyAtPercentile(latencyPercentile)),
			results.Duration(scenario.Name),
			results.Rps(scenario.Name),
			results.ThresholdsPassed(scenario.Name),
		})
	}

	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	output.WriteString(fmt.Sprintf("Total duration: %v\n", results.TotalDuration()))

	if success {
		output.WriteString("OK")
	} else {
		output.WriteString("Failed")
	}

	return output.String()
}
