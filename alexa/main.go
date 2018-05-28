package main

import (
	"fmt"
	"os"

	"github.com/danielcondemarin/go-ci-commander/teamcity"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var applications = map[string]interface{}{
	"/echo/cicommander": alexa.EchoApplication{
		AppID:    getEnvVar("ALEXA_APP_ID", true),
		OnIntent: onIntentHandler,
		OnLaunch: onLaunchHandler,
	},
}

func main() {
	alexa.Run(applications, "3000")
}

func getEnvVar(key string, throw bool) string {
	value := os.Getenv(key)

	if len(value) <= 0 && throw {
		panic(fmt.Sprintf("Missing env. variable: %s", key))
	}

	return value
}

func onLaunchHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeechSSML("<speak>Welcome to Commander c.i.</speak>")
}

func onIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	switch echoReq.GetIntentName() {
	case "TriggerBuild":
		buildTypeID, _ := echoReq.GetSlotValue("buildType")
		teamcityURL := getEnvVar("TEAMCITY_URL", true)
		teamcityUser := getEnvVar("TEAMCITY_USER", true)
		teamcityPass := getEnvVar("TEAMCITY_PASS", true)

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
