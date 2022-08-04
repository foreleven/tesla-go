package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"tesla-go/graph/generated"
	"tesla-go/tesla"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, token string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Vehicles is the resolver for the vehicles field.
func (r *queryResolver) Vehicles(ctx context.Context) ([]*tesla.Vehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *vehicleResolver) ID(ctx context.Context, obj *tesla.Vehicle) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// VehicleID is the resolver for the vehicleID field.
func (r *vehicleResolver) VehicleID(ctx context.Context, obj *tesla.Vehicle) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Color is the resolver for the color field.
func (r *vehicleResolver) Color(ctx context.Context, obj *tesla.Vehicle) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
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
