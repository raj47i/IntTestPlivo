package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raj47i/IntTestPlivo/models"
)

func smsOutbound(c *gin.Context) {
	var sms models.SMS
	if err := sms.Parse(c.PostForm("to"), c.PostForm("from"), c.PostForm("text")); err != nil {
		http422(c, err)
		return
	}
	accoundID := c.MustGet("AccountID").(uint)
	var pn models.PhoneNumber
	if loaded, err := pn.LoadByNumberAndAccountID(sms.From, accoundID); err != nil {
		http500(c, errors.New("unknown failure"))
		return
	} else if !loaded {
		http422(c, errors.New("from parameter not found"))
		return
	}

	if blocked, err := sms.IsBlocked(); err != nil {
		http500(c, fmt.Errorf("unknown failure"))
		return
	} else if blocked {
		http422(c, fmt.Errorf("sms from %s to %s blocked by STOP request", sms.From, sms.To))
		return
	}

	if exceeded, err := sms.DayLimit(); err != nil {
		http500(c, fmt.Errorf("unknown failure"))
		return
	} else if exceeded {
		http403(c, fmt.Errorf("limit reached for from %s", sms.From))
		return
	}

	http200(c, "outbound sms ok")
}
