package discovery

// based on examples provided by Travis Jeffery in his book Distributed Services with Go

import(
	"net"
	"go.uber.org/zap"
	"github.com/hashicorp/serf"
)

type Membership struct {
	Config
	handler Handler
	serf *serf.Serf
	events chan serf.Event
	logger *zap.logger
}

func New(handler Handler, config Config) (*Membership, error) {
	c := &Membership{
		Config: config,
		handler: handler,
		logger: zap.L().Named("membership")
	}

	if err := c.setupSerf(); err != nil {
		return nil, err
	}

	return c, nil
}