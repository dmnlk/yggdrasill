package gomadare

import "github.com/mrjones/oauth"

type Client struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
}

const (
	RequestTokenUrl   string = "http://api.twitter.com/oauth/request_token"
	AuthorizeTokenUrl string = "https://api.twitter.com/oauth/authorize"
	AccessTokenUrl    string = "https://api.twitter.com/oauth/access_token"
)


// initialize client, need consumer key, consumersecret, accesstoken, accesstokensecret,from twitter dev
func NewClient(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Client {
	client := new(Client)
	client.consumer = oauth.NewConsumer(consumerKey, consumerSecret, oauth.ServiceProvider{
		RequestTokenUrl:   RequestTokenUrl,
		AuthorizeTokenUrl: AuthorizeTokenUrl,
		AccessTokenUrl:    AccessTokenUrl,
	})
	client.accessToken = &oauth.AccessToken{accessToken, accessTokenSecret, nil}
	return client
}
