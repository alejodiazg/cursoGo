package service

import (
	"fmt"
	"github.com/alejodiazg/cursoGo/src/domain"
	"time"
)


type TweetManager struct {
	users map[string]*domain.User
	tweets []*domain.Tweet
	userTweets map[string][]*domain.Tweet
}
var users  = make(map[string]*domain.User)
var tweets []*domain.Tweet
var userTweets = make(map[string][]*domain.Tweet)

func NewTweetManager() TweetManager {
	return  TweetManager{
		users  : make(map[string]*domain.User),
		tweets : make([]*domain.Tweet,0) ,
		userTweets : make(map[string][]*domain.Tweet),
	}
}

func (tm *TweetManager) PublishTweet(toPublish *domain.Tweet)  (int,error) {
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
		tm.tweets = append(tm.tweets, toPublish)
		_ , ok := tm.userTweets[toPublish.User]
		if !ok {
			tm.userTweets[toPublish.User] = make([]*domain.Tweet, 0)
		}
		tm.userTweets[toPublish.User] = append(tm.userTweets[toPublish.User], toPublish)
	}

	return toPublish.Id , err
}

func (tm TweetManager) GetTweet() *domain.Tweet{
	return tm.tweets[len(tm.tweets)-1]
}


func (tm TweetManager) GetTweets() []*domain.Tweet {
	return tm.tweets
}

func (tm TweetManager) GetTweetById(id int) *domain.Tweet{
	for _, x := range tm.tweets {
		if x.Id == id {
			return x
		}
	}
	return nil
}

func (tm TweetManager) CountTweetsByUser(user string) int {
	list  := tm.userTweets[user]
	return len(list)
}

func (tm TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	list , ok := tm.userTweets[user]
	if !ok {
		list = make([]*domain.Tweet, 0)
	}
	return list
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

