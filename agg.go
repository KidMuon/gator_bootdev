package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kidmuon/gator_bootdev/internal/database"
)

func handlerAgg(state *State, command Command) error {
	if len(command.args) < 1 {
		return fmt.Errorf("agg requires a time duration string to set the time between requests")
	}

	time_between_reqs, err := time.ParseDuration(command.args[0])
	if err != nil {
		return fmt.Errorf("cannot parse duration from argument %s: %v", command.args[0], err)
	}

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, bytes.NewBuffer([]byte{}))
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Add("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	var rssFeed RSSFeed
	if err = xml.Unmarshal(data, &rssFeed); err != nil {
		return &RSSFeed{}, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for _, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}

func printRSSItems(rssFeed *RSSFeed) {
	for _, item := range rssFeed.Channel.Item {
		fmt.Println("\n\n--------------------------------------")
		fmt.Printf("Title: %s\n\n", item.Title)
		fmt.Printf("Link: %s\n\n", item.Link)
		fmt.Printf("Description: %s\n\n", item.Description)
		fmt.Printf("Publication Date: %s\n", item.PubDate)
		fmt.Println("--------------------------------------")
	}
}

func scrapeFeeds(state *State) error {
	feed, err := state.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("fetching next feed failed: %v", err)
	}

	markFeed := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            feed.ID,
	}
	_, err = state.db.MarkFeedFetched(context.Background(), markFeed)
	if err != nil {
		return fmt.Errorf("marking feed update failed: %v", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}

	saveFeedPosts(state, rssFeed, feed)

	return nil
}

func saveFeedPosts(state *State, rssFeed *RSSFeed, feed database.Feed) error {
	var layout string = "12-31-2019 8:00:02"
	var url_error string = "pq: duplicate key value violates unique constraint \"posts_url_key\""

	for _, post := range rssFeed.Channel.Item {
		sqlTitle := post.Title

		var sqlLink sql.NullString
		if post.Link == "" {
			sqlLink = sql.NullString{Valid: false}
		} else {
			sqlLink = sql.NullString{String: post.Link, Valid: true}
		}

		var sqlDescription sql.NullString
		if post.Description == "" {
			sqlDescription = sql.NullString{Valid: false}
		} else {
			sqlDescription = sql.NullString{String: post.Description, Valid: true}
		}

		var sqlPubTime sql.NullTime
		if post.PubDate == "" {
			sqlPubTime = sql.NullTime{Valid: false}
		} else {
			parsedTime, _ := time.Parse(layout, post.PubDate)
			sqlPubTime = sql.NullTime{Time: parsedTime, Valid: true}
		}

		postToCreate := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sqlTitle,
			Url:         sqlLink,
			Description: sqlDescription,
			PublishedAt: sqlPubTime,
			FeedID:      feed.ID,
		}

		_, err := state.db.CreatePost(context.Background(), postToCreate)
		if err != nil && err.Error() != url_error {
			fmt.Printf("post {%s, %s} creation failed: %v\n", post.Title, feed.Name, err)
		}
	}

	return nil
}
