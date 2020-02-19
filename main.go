package main

import (
	"net/url"
	// "os"
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
)

var (
	consumerKey       = "fOm54y9kep6yDIeYQc60oE6x2"
	consumerSecret    = "CTHtwLktuCMJRa2utU4UBuDGNDvk5VqgM9f95MtAYH3Xm2Lkhh"
	accessToken       = "1216805040419147776-M6rudtnfKW2Oi0lzgmBiC6DyK2Kbl1"
	accessTokenSecret = "GetxXsVnfhNxsN2fH8XMARWl5Au5MmRl4URCwWA7wNV9x"
)

// var (
// 	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
// 	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
// 	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
// 	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
// )

// func getenv(name string) string {
// 	v := os.Getenv(name)
// 	if v == "" {
// 		panic("missing required environment variable " + name)
// 	}
// 	return v
// }

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{
			"@IAmBoredBot I am bored", 
			"@IAmBoredBot suggestion", 
			"@IAmBoredBot what should I do", 
			},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("received unexpected value of type %T and not a type of anaconda.Tweet", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}

		_, err := api.Retweet(t.Id, false)
		
		if err != nil {
			log.Errorf("could not retweet %d: %v", t.Id, err)
			continue
		}
		log.Infof("retweeted %d", t.Id)
	}
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { 
	log.Error(args...)
}

func (log *logger) Criticalf(format string, args ...interface{}) { 
	log.Errorf(format, args...) 
}

func (log *logger) Notice(args ...interface{})                   { 
	log.Info(args...) 
}

func (log *logger) Noticef(format string, args ...interface{})   { 
	log.Infof(format, args...) 
}