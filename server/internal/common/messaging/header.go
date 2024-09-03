package messaging

import (
	"time"

	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	events "github.com/subscribeddotdev/subscribed/server/events/go"
)

func NewHeader(eventName, serviceName string) *events.Header {
	return &events.Header{
		Id:            ulid.Make().String(),
		Name:          eventName,
		CorrelationId: "", // TODO: set the corr id via middleware
		PublisherName: serviceName,
		PublishedAt:   timestamppb.New(time.Now()),
	}
}
