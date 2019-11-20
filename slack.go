package main

import (
	"github.com/nlopes/slack"
	"strings"
)

type Client struct {
	*slack.Client
}

type Channel struct {
	ID string
	Name string `json:"name"`
}

func getClient(token string) Client {
	client := slack.New(token)

	return Client{client}
}

func (c *Client) getChannelByName(name string) (Channel, error) {

	channels, err := c.GetChannels(false)
	for _, channel := range channels {
		if strings.ToLower(channel.Name) == strings.ToLower(name) {
			return Channel{channel.ID, channel.Name}, nil
		}
	}

	groups, err := c.GetGroups(false)

	for _, group := range groups {
		if strings.ToLower(group.Name) == strings.ToLower(name) {
			return Channel{group.ID, group.Name}, nil
		}
	}

	return Channel{}, err
}

func (c *Client) sendMessage(channelId, message string, requestData RequestData) error {
	link := slack.Attachment{
		Text: requestData.Attributes.Url,
	}

	message = strings.Replace(message, "_user_", requestData.User.Name, 1)
	message = strings.Replace(message, "_project_", requestData.Project.Name, 1)

	_, _, _, err := c.SendMessage(channelId, slack.MsgOptionText("<!here> "+message, false), slack.MsgOptionAsUser(true), slack.MsgOptionAttachments(link))

	return err
}