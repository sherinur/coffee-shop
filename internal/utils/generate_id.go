package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func GenerateNewID(items []string, prefix string) string {
	if len(items) == 0 {
		return fmt.Sprintf("%s1", prefix)
	}

	var maxID int
	re := regexp.MustCompile(fmt.Sprintf(`%s(\d+)`, prefix))

	for _, item := range items {
		matches := re.FindStringSubmatch(item)
		if len(matches) > 1 {
			id, err := strconv.Atoi(matches[1])
			if err == nil && id > maxID {
				maxID = id
			}
		}
	}

	return fmt.Sprintf("%s%d", prefix, maxID+1)
}
