package service

import (
	"fmt"
	"github.com/alejodiazg/cursoGo/src/domain"
	"time"
)

var users  = make(map[string]*domain.User)
var tweets []*domain.Tweet

func InitializeService() {
	tweets = nil
	users = make(map[string]*domain.User)
}

func PublishTweet(toPublish *domain.Tweet)  (int,error) {
	var err error = nil
	if toPublish.User == "" {
		err =  fmt.Errorf("user is required")
	} else if toPublish.Text == "" {
		err =  fmt.Errorf("text is required")
	} else if len(toPublish.Text) > 140 {
		err =  fmt.Errorf("text exceeds 140 characters")
	}

	if err == nil {
		toPublish.Id = int(time.Now().Unix())
		tweets = append(tweets, toPublish)
	}

	return toPublish.Id , err
}

func GetTweet() *domain.Tweet{
	return tweets[len(tweets)-1]
}


func GetTweets() []*domain.Tweet {
	return tweets
}

func GetTweetById(id int) *domain.Tweet{
	for _, x := range tweets {
		if x.Id == id {
			return x
		}
	}
	return nil
}

func CountTweetsByUser(user string) int {
	var i = 0
	for _, x := range tweets {
		if x.User == user {
			i++
		}
	}
	return i
}

func RegisterUser(username, name, email, password string) error {
	user := domain.NewUser(username, name, email, password)
	_, ok := users[user.Email]
	if ok {
		return fmt.Errorf("text exceeds 140 characters")
	}

	users[user.Email] = user
	return nil
}

