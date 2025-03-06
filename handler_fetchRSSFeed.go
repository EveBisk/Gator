package main

import (
	"fmt"
)

func handlerFetchFeed(s *state, cmd command) error {
	// if len(cmd.args) != 1 {
	// 	return fmt.Errorf("usage: %v <name>", cmd.name)
	// }

	// feedUrl := cmd.args[0]
	feedUrl := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(s.ctx, feedUrl)

	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	printRSSFeed(*feed)
	return nil
}
