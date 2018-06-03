package main

import (
	"fmt"

	alexa "github.com/danielcondemarin/go-alexa/skillserver"
	"github.com/danielcondemarin/go-ci-commander/env"
	"github.com/danielcondemarin/go-ci-commander/teamcity"
)

var applications = map[string]interface{}{
	"/echo/teamcity": alexa.EchoApplication{
		AppID:    env.GetEnvVar("ALEXA_APP_ID", true),
		OnIntent: onIntentHandler,
		OnLaunch: onLaunchHandler,
	},
}

func main() {
	alexa.Run(applications, "3000")
}

func onLaunchHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Welcome to TeamCity")
}

func onIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	switch echoReq.GetIntentName() {
	case "TriggerBuild":
		slot, _ := echoReq.GetSlot("buildType")
		slotValue := slot.Value

		// check for original value in case a synonym was used
		ra := slot.Resolutions.ResolutionsPerAuthority
		if len(ra) > 0 {
			slotValue = ra[0].Values[0]["value"].Name
		}

		teamcityURL := env.GetEnvVar("TEAMCITY_URL", true)
		teamcityUser := env.GetEnvVar("TEAMCITY_USER", true)
		teamcityPass := env.GetEnvVar("TEAMCITY_PASS", true)

		client := teamcity.NewClient(teamcityURL, teamcity.NewBasicHTTPAuthorizer(teamcityUser, teamcityPass))

		r := teamcity.NewBuildRequest(slotValue)
		_, err := client.TriggerBuild(r)

		if err != nil {
			echoResp.OutputSpeech("There was a problem trying to run that build configuration")
		} else {
			echoResp.OutputSpeech(fmt.Sprintf("TeamCity Build %s triggered", slotValue))
		}
	}
}
