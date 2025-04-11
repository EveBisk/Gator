package main

import (
	"fmt"
	"time"
)

func handlerFetchFeed(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	time_between_reqs := cmd.args[0]
	time_duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("couldn't parse time duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s", time_between_reqs)

	ticker := time.NewTicker(time_duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
