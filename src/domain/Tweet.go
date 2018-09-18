package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetString() string
	GetUser()   string
	GetDate()   *time.Time
	GetText()   string
	GetId() int
	SetId(int)
	PrintableTweet() string
}

type TextTweet struct {
	user string
	text string
	date *time.Time
	id   int
}

func NewTextTweet(user string , text string) *TextTweet {
	date := time.Now().UTC()
	return  &TextTweet{user, text, &date, 0}
}

func (tweet TextTweet) GetString() string {
	return fmt.Sprintf("@%s: %s", tweet.user, tweet.text)
}

func (tweet TextTweet) GetText() string {
	return tweet.text
}

func (tweet TextTweet) GetUser() string {
	return tweet.text
}

func (tweet TextTweet) GetDate() *time.Time {
	return tweet.date
}

func (tweet TextTweet) GetId() int {
	return tweet.id
}

func (tweet *TextTweet) SetId(id int)  {
	tweet.id = id
}

func (tweet TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.user, tweet.text)
}

func (tweet TextTweet) String() string {
	return fmt.Sprintf("@%s: %s", tweet.user, tweet.text)
}

//image tweet
type ImageTweet struct {
	TextTweet
	image string
}

func NewImageTweet(user string , text string, image string) *ImageTweet{
	date := time.Now().UTC()
	return &ImageTweet{
		TextTweet:  TextTweet{user, text, &date, 0},
		image: image,
	}
}

func (tweet ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", tweet.user, tweet.text, tweet.image)
}

//quote  tweet
type QuoteTweet struct {
	TextTweet
	quotedText string
}

func NewQuoteTweet(user string , text string, quoted Tweet) *QuoteTweet {
	date := time.Now().UTC()
	return &QuoteTweet{
		TextTweet:  TextTweet{user, text, &date, 0},
		quotedText: quoted.PrintableTweet(),
	}
}

func (tweet QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"%s\"", tweet.user, tweet.text, tweet.quotedText)
}
