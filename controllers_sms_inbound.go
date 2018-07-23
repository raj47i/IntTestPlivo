package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/raj47i/IntTestPlivo/models"
)

func smsInbound(c *gin.Context) {
	var sms models.SMS
	if err := sms.Parse(c.PostForm("to"), c.PostForm("from"), c.PostForm("text")); err != nil {
		http422(c, err)
		return
	}
	accoundID := c.MustGet("AccountID").(uint)
	var pn models.PhoneNumber
	if loaded, err := pn.LoadByNumberAndAccountID(sms.To, accoundID); err != nil {
		http500(c, errors.New("unknown failure"))
		return
	} else if !loaded {
		http422(c, errors.New("to parameter not found"))
		return
	}
	if sms.IsBlockCommand() {
		if err := sms.Block(); err != nil {
			http500(c, errors.New("unknown failure"))
			return
		}
	}
	http200(c, "inbound sms ok")
}
