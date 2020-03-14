package ssh

const (
	agentRequestType = "auth-agent-req@openssh.com"
)

// contextKeyAgentRequest is an internal context key for storing if
// the client requested agent forwarding
var contextKeyAgentRequest = &contextKey{"auth-agent-req"}

// SetAgentRequested sets up the session context so that AgentRequested
// returns true.
func SetAgentRequested(ctx Context) {
	ctx.SetValue(contextKeyAgentRequest, true)
}
