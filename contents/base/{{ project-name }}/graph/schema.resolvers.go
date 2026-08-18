package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"{{ module_path }}/graph/model"
	"{{ module_path }}/internal/repository"
)

// Create{{ EntityName }} is the resolver for the create{{ EntityName }} field.
func (r *mutationResolver) Create{{ EntityName }}(ctx context.Context, displayName string) (*model.{{ EntityName }}, error) {
	e, err := r.Store.Create(ctx, displayName)
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

// Update{{ EntityName }} is the resolver for the update{{ EntityName }} field.
func (r *mutationResolver) Update{{ EntityName }}(ctx context.Context, id string, displayName string) (*model.{{ EntityName }}, error) {
	e, err := r.Store.Update(ctx, id, displayName)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

// Delete{{ EntityName }} is the resolver for the delete{{ EntityName }} field.
func (r *mutationResolver) Delete{{ EntityName }}(ctx context.Context, id string) (bool, error) {
	if err := r.Store.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// {{ EntityName }} is the resolver for the {{ entityName }} field.
func (r *queryResolver) {{ EntityName }}(ctx context.Context, id string) (*model.{{ EntityName }}, error) {
	e, err := r.Store.Get(ctx, id)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

// {{ EntityName }}s is the resolver for the {{ entityName }}s field.
func (r *queryResolver) {{ EntityName }}s(ctx context.Context) ([]*model.{{ EntityName }}, error) {
	items, err := r.Store.List(ctx)
	if err != nil {
		return nil, err
	}
	out := []*model.{{ EntityName }}{}
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
