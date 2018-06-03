# Trigger your Teamcity builds from Alexa!

## Requirements:
* `ask-cli` package https://www.npmjs.com/package/ask-cli
	* `npm i -g ask-cli`
* `ngrok` and `golang` if using the local Alexa Server instead of a Lambda in AWS

## If the Teamcity Server is not public on the internet:

You're covered, the repo includes an Alexa Skill Server built in Go which you can use instead of AWS Lambda.

Steps:

*	`ngrok http 3000` *# Paste the HTTPs URL onto `apis.custom.endpoint.uri` in `alexa-teamcity/skill.json`, leave the path `echo/teamcity` intact*
*	`cd go-ci-commander/alexa-teamcity`
*	`ask init` *# Make sure you choose a profile which has sufficient permissions, more [here](https://developer.amazon.com/docs/smapi/ask-cli-command-reference.html#init-command).*
*	`ask deploy` *# After completed copy the `skill_id` in `./ask/config`, you'll use it below*
*	`go build *.go`
*	`ALEXA_APP_ID="{skill_id}" TEAMCITY_URL="http://{teamcity_host}:{port}" TEAMCITY_USER="{user}" TEAMCITY_PASS="{pwd}" ./main`

## Testing it ...

*	Add your Build configuration IDs to `types.values` in `models/en-GB.json`. You can also use synonyms for your build IDs so they are easier to pronounce etc.
*   Redeploy the skill model, `ask deploy` 
*	`ask simulate --text "ask teamcity to trigger {myBuildType}" --locale "en-GB"` *# or simply use the Alexa Test simulator, Echoism etc.*

## If the Teamcity Server is publicly available:

In progress, working on a Lambda alternative to the local Alexa Skill Server

