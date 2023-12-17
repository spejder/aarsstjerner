package ms

import (
	"reflect"

	"bitbucket.org/long174/go-odoo"
)

// Client provides high and low level functions to interact with Medlemsservice.
type Client struct {
	odoo.Client
}

// UID of the authenticated client.
func (c *Client) UID() int64 {
	return reflect.ValueOf(c.Client).FieldByName("uid").Int()
}
