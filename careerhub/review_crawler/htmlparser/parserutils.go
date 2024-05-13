package htmlparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jae2274/goutils/terr"
)

func ParseReviewScore(scoreStr string) (int32, error) {
	if !scoreRegex.MatchString(scoreStr) {
		return 0, terr.New(fmt.Sprintf("invalid score format. expected: 0.0~5.0, got: %s", scoreStr))
	}

	splits := strings.Split(scoreStr, ".")

	tenDigit, err := strconv.Atoi(splits[0])
	if err != nil {
		return 0, err
	}
	oneDigit, err := strconv.Atoi(splits[1])
	if err != nil {
		return 0, err
	}

	return int32(tenDigit*10 + oneDigit), nil
}

func ParseReviewCount(countStr string) (int32, error) {
	countStr = strings.ReplaceAll(countStr, ",", "")
	re := regexp.MustCompile(`\d+`)
	countStr = re.FindString(countStr)
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}

	return int32(count), nil
}
