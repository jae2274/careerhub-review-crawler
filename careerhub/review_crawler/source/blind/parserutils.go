package blind

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

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

func ParseReviewUser(str string) (string, string, time.Time, error) {
	dateParts := strings.Split(str, "-")

	if len(dateParts) != 2 {
		return "", "", time.Time{}, terr.New(fmt.Sprintf("invalid review user format. got: %s", str))
	}

	dateStr := strings.TrimSpace(dateParts[1])
	date, err := time.Parse("2006.01.02", dateStr)

	if err != nil {
		return "", "", time.Time{}, err
	}

	userParts := strings.Split(dateParts[0], "·")

	if len(userParts) < 3 {
		return "", "", time.Time{}, terr.New(fmt.Sprintf("invalid review user format. got: %s", str))
	}

	userId := strings.TrimSpace(userParts[1])

	jobType := strings.TrimSpace(strings.Join(userParts[2:], "·"))

	return userId, jobType, date, nil
}
