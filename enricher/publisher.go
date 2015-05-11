package enricher

import(
	sp "github.com/wunderlist/snowblower/snowplow"
)

type Publisher interface {
	publish(event *sp.Event)
}