package ticker

import (
	"fmt"
	"strings"
)

// GenYahooTickerFromTipRankTicker gen yahoo ticker from tip rank ticker
func GenYahooTickerFromTipRankTicker(tiprankTicker string) string {
	parts := strings.Split(tiprankTicker, ":")

	if len(parts) > 1 {
		if parts[0] == "TSE" {
			return fmt.Sprintf("%s.TO", strings.Replace(parts[1], ".", "-", -1))
		}
	}

	return fmt.Sprintf("%s.TO", tiprankTicker)
}
