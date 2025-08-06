package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	goonvif "github.com/guchengod/go-onvif"
	"github.com/guchengod/go-onvif/device"
	sdk "github.com/guchengod/go-onvif/sdk/device"
	"github.com/guchengod/go-onvif/xsd/onvif"
)

const (
	login    = "admin"
	password = "hzmt2025"
)

func main() {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.20.73:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	//Preparing commands
	systemDateAndTyme := device.GetSystemDateAndTime{}
	getCapabilities := device.GetCapabilities{Category: "All"}
	createUser := device.CreateUsers{
		User: onvif.User{
			Username:  "TestUser",
			Password:  "TestPassword",
			UserLevel: "User",
		},
	}

	//Commands execution
	systemDateAndTymeResponse, err := sdk.Call_GetSystemDateAndTime(ctx, dev, systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse)
	}
	getCapabilitiesResponse, err := sdk.Call_GetCapabilities(ctx, dev, getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}

	createUserResponse, err := sdk.Call_CreateUsers(ctx, dev, createUser)
	if err != nil {
		log.Println(err)
	} else {
		// You could use https://github.com/guchengod/go-onvif/gosoap for pretty printing response
		fmt.Println(createUserResponse)
	}

}
