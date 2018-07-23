package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/raj47i/IntTestPlivo/config"
)

// SMS represents a text message, with recipient and sender
type SMS struct {
	To   string
	From string
	Text string
}

// Parse will parse and validate SMS from input
func (sms *SMS) Parse(to string, from string, text string) error {
	if len(to) < 1 {
		return errors.New("to is missing")
	}
	if len(from) < 1 {
		return errors.New("from is missing")
	}
	if len(text) < 1 {
		return errors.New("text is missing")
	}

	rePhone := regexp.MustCompile("^[0-9]{6,16}$")
	if !rePhone.MatchString(to) {
		return errors.New("to is invalid")
	}
	if !rePhone.MatchString(from) {
		return errors.New("from is invalid")
	}
	if len(text) > 120 {
		return errors.New("text is invalid")
	}
	sms.To = to
	sms.From = from
	sms.Text = text
	return nil
}

func (sms *SMS) stopKey() string {
	return fmt.Sprintf("STOP:%s:%s", sms.From, sms.To)
}
func (sms *SMS) daylimitKey() string {
	return fmt.Sprintf("LAST24:%s", sms.From)
}

// IsBlocked checks if the message is blocked
func (sms *SMS) IsBlocked() (bool, error) {
	c := config.GetCache()
	defer c.Close()
	res, err := redis.Bool(c.Do("GET", sms.stopKey()))
	if err == redis.ErrNil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return res, nil
}

// Block blocks messages for 4 hours
func (sms *SMS) Block() error {
	c := config.GetCache()
	defer c.Close()
	_, err := c.Do("SETEX", sms.stopKey(), 14400, true) // 4 hours = 14400 seconds
	return err
}

// IsBlockCommand checks if the sms is a BLOCKING instruction
func (sms *SMS) IsBlockCommand() bool {
	return "STOP" == strings.TrimRight(sms.Text, "\r\n")
}

// DayLimit checks  if number of messages in last 24 hours exceeded 50 & counts
func (sms *SMS) DayLimit() (bool, error) {
	c := config.GetCache()
	defer c.Close()

	if exists, err := redis.Bool(c.Do("EXISTS", sms.daylimitKey())); err != nil {
		return false, err
	} else if exists {
		if res, err := c.Do("INCR", sms.daylimitKey()); err != nil {
			return false, err
		} else if val, err := redis.Int(res, nil); err != nil {
			return false, err
		} else if val > 50 {
			return true, nil
		}
	} else if _, err := c.Do("SETEX", sms.daylimitKey(), 86400, 1); err != nil { //24 hours = 86400 seconds
		return false, err
	}
	return false, nil
}
