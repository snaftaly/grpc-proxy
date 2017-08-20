// Copyright 2017 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package proxy

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// StreamDirector returns a gRPC ClientConn to be used to forward the call to, and a Context that can
// be altered by the director and used later on.
//
//
// The presence of the `Context` allows for rich filtering, e.g. based on Metadata (headers).
// If no handling is meant to be done, a `codes.NotImplemented` gRPC error should be returned.
//
// It is worth noting that the StreamDirector will be fired *after* all server-side stream interceptors
// are invoked. So decisions around authorization, monitoring etc. are better to be handled there.
//
// See the rather rich example.

type ConnectionWithContext struct {
	ClientConnection *grpc.ClientConn
	Ctx              context.Context
}

type StreamDirector func(ctx context.Context, fullMethodName string) (*ConnectionWithContext, error)
