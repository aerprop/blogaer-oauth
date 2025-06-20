package enum

// Consumers
const (
	googleC  string = "GoogleConsumer"
	githubC string = "GithubConsumer"
)

var Consumer = struct {
	Google string
	Github string
}{
	Google: googleC,
	Github: githubC,
}

var AllConsumers = []string{
	Consumer.Google,
	Consumer.Github,
}

// exchanges
const (
	rpcExchName string = "RpcExchange"
	rpcExchType string = "direct"
)

var RpcExchange = struct {
	Name string
	Type string
}{
	Name: rpcExchName,
	Type: rpcExchType,
}

// queues
const (
	googleQ string = "OauthGoogleQueue"
	githubQ string = "OauthGithubQueue"
)

var QueueName = struct {
	Google string
	Github string
}{
	Google: googleQ,
	Github: githubQ,
}

var AllQueueNames = []string{
	QueueName.Google,
	QueueName.Github,
}

// route keys
const (
	googleRK string = "oauth.google.key"
	githubRK string = "oauth.github.key"
)

var RoutingKey = struct {
	GoogleRK string
	GithubRK string
}{
	GoogleRK: googleRK,
	GithubRK: githubRK,
}

var AllRouteKeys = []string{
	RoutingKey.GoogleRK,
	RoutingKey.GithubRK,
}

// oauth2 providers
const (
	google string = "google"
	github string = "github"
)

var OauthProvider = struct {
	Google string
	Github string
}{
	Google: google,
	Github: github,
}

var AllOauthProviders = []string{
	OauthProvider.Google,
	OauthProvider.Github,
}
