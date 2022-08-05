package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"
	"tesla-go/core"
	"tesla-go/graph/generated"
	"tesla-go/tesla"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, token string) (string, error) {
	auth, err := core.Create(token)
	if err != nil {
		return "", errors.New("bad token")
	} else {
		return auth.Token(), nil
	}
}

// Vehicles is the resolver for the vehicles field.
func (r *queryResolver) Vehicles(ctx context.Context) ([]*tesla.Vehicle, error) {
	c := core.ForContext(ctx)
	c.HasAuthorized()
	client := c.GetClient()
	return client.Vehicles()
}

// ID is the resolver for the id field.
func (r *vehicleResolver) ID(ctx context.Context, obj *tesla.Vehicle) (string, error) {
	return string(obj.ID), nil
}

// VehicleID is the resolver for the vehicleID field.
func (r *vehicleResolver) VehicleID(ctx context.Context, obj *tesla.Vehicle) (string, error) {
	return strconv.FormatInt(obj.VehicleID, 10), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Vehicle returns generated.VehicleResolver implementation.
func (r *Resolver) Vehicle() generated.VehicleResolver { return &vehicleResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type vehicleResolver struct{ *Resolver }
