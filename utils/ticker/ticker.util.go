package ticker

import (
	"fmt"
	"strings"

	"github.com/lenoobz/aws-tiprank-norm-dividend/consts"
)

// GenYahooTickerFromTipRankTicker gen yahoo ticker from tip rank ticker
func GenYahooTickerFromTipRankTicker(tiprankTicker string) string {
	specialTicker := getSpecialYahooTicker(tiprankTicker)
	if specialTicker != "" {
		return specialTicker
	}

	parts := strings.Split(tiprankTicker, ":")

	if len(parts) > 1 {
		if parts[0] == "TSE" {
			return fmt.Sprintf("%s.TO", strings.Replace(parts[1], ".", "-", -1))
		}
	}

	return tiprankTicker
}

// getSpecialYahooTicker gets special yahoo ticker from a given tiprank ticker
func getSpecialYahooTicker(tiprankTicker string) string {
	for _, ticker := range consts.TipRankToYahooTickers {
		if strings.EqualFold(strings.ToUpper(ticker.TipRank), strings.ToUpper(tiprankTicker)) {
			return ticker.Yahoo
		}
	}

	return ""
}
