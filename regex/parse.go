package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := []byte(`openebs_zfs_stats_reject_request_count 16`)
	re := regexp.MustCompile(`openebs_zfs_stats_reject_request_count \d+`)
	result := re.FindStringSubmatch(string(s))
	fmt.Println(result)
}
