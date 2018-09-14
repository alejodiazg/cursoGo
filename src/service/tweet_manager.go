package service

import (
	"fmt"
	"github.com/alejodiazg/cursoGo/src/domain"
)

var tweet *domain.Tweet

func PublishTweet(toPublish *domain.Tweet) (err error){
	if toPublish.User == "" {
		return fmt.Errorf("user is required")
	}

	if toPublish.Text == "" {
		return fmt.Errorf("text is required")
	}

	if toPublish.Text == "" {
		return fmt.Errorf("text is required")
	}

	if len(toPublish.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}

	tweet = toPublish
	return nil
}

func GetTweet() *domain.Tweet{
	return tweet
}