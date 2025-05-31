package enum

// channels
const (
	consumer  string = "OauthRpcChannel"
	publisher string = "OauthPubChannel"
)

var Channel = struct {
	Consumer  string
	Publisher string
}{
	Consumer:  consumer,
	Publisher: publisher,
}

var AllChannels = []string{
	Channel.Consumer,
	Channel.Publisher,
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
	GoogleQ string
	GithubQ string
}{
	GoogleQ: googleQ,
	GithubQ: githubQ,
}

var AllQueueNames = []string{
	QueueName.GoogleQ,
	QueueName.GithubQ,
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
