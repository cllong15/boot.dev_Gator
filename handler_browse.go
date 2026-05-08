package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	var ok error
	if len(cmd.Args) == 1 {
		limit, ok = strconv.Atoi(cmd.Args[0])
		if ok != nil {
			return fmt.Errorf("could not parse limit: %w", ok)
		}
	} else {
		limit = 2
	}
	posts, ok := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if ok != nil {
		return fmt.Errorf("could not get posts for %s: %w", user.Name, ok)
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	return nil
}
