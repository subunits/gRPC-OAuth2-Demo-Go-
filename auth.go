package main

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

// Simple JWT-like check
func validateJWT(token string) bool {
	// Replace with real JWT verification
	return token == "valid-demo-token"
}

// Unary interceptor
func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeaders := md["authorization"]
	if len(authHeaders) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}

	token := strings.TrimPrefix(authHeaders[0], "Bearer ")
	if !validateJWT(token) {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	// Token valid, continue
	return handler(ctx, req)
}
