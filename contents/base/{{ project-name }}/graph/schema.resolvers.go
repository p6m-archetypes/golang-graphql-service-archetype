package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"{{ module_path }}/graph/model"
	"{{ module_path }}/internal/repository"
)

// Create{{ PrefixName }} is the resolver for the create{{ PrefixName }} field.
func (r *mutationResolver) Create{{ PrefixName }}(ctx context.Context, displayName string) (*model.{{ PrefixName }}, error) {
	e, err := r.Store.Create(ctx, displayName)
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

// Update{{ PrefixName }} is the resolver for the update{{ PrefixName }} field.
func (r *mutationResolver) Update{{ PrefixName }}(ctx context.Context, id string, displayName string) (*model.{{ PrefixName }}, error) {
	e, err := r.Store.Update(ctx, id, displayName)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

// Delete{{ PrefixName }} is the resolver for the delete{{ PrefixName }} field.
func (r *mutationResolver) Delete{{ PrefixName }}(ctx context.Context, id string) (bool, error) {
	if err := r.Store.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// {{ PrefixName }} is the resolver for the {{ prefix_name | camel_case }} field.
func (r *queryResolver) {{ PrefixName }}(ctx context.Context, id string) (*model.{{ PrefixName }}, error) {
	e, err := r.Store.Get(ctx, id)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

// {{ PrefixName }}s is the resolver for the {{ prefix_name | camel_case }}s field.
func (r *queryResolver) {{ PrefixName }}s(ctx context.Context) ([]*model.{{ PrefixName }}, error) {
	items, err := r.Store.List(ctx)
	if err != nil {
		return nil, err
	}
	out := []*model.{{ PrefixName }}{}
	for _, e := range items {
		out = append(out, toModel(e))
	}
	return out, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
