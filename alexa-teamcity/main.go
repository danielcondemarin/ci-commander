package main

import (
	"fmt"

	"github.com/danielcondemarin/go-ci-commander/env"
	"github.com/danielcondemarin/go-ci-commander/teamcity"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
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
		buildTypeID, _ := echoReq.GetSlotValue("buildType")
		teamcityURL := env.GetEnvVar("TEAMCITY_URL", true)
		teamcityUser := env.GetEnvVar("TEAMCITY_USER", true)
		teamcityPass := env.GetEnvVar("TEAMCITY_PASS", true)

		client := teamcity.NewClient(teamcityURL, teamcity.NewBasicHTTPAuthorizer(teamcityUser, teamcityPass))

		r := teamcity.NewBuildRequest(buildTypeID)
		_, err := client.TriggerBuild(r)

		if err != nil {
			panic(err)
		} else {
			echoResp.OutputSpeech(fmt.Sprintf("TeamCity Build %s triggered", buildTypeID))
		}
	}
}
