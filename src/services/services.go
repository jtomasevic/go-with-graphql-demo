package services

import seas "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas"

// Services: instance of this struct keep references to all serivces. In our case for now we
// have only one service, but usually we'll have more than one, and then this struct become more meaningful.
// This struct is passed to resolvers, so they can be free to use service through one point
type Services struct {
	SevenSeasService seas.Service
}
