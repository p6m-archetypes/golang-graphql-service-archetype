package graph

import (
	"{{ module_path }}/graph/model"
	"{{ module_path }}/internal/repository"
)

// toModel maps the repository entity to its GraphQL model.
func toModel(e repository.{{ EntityName }}) *model.{{ EntityName }} {
	return &model.{{ EntityName }}{ID: e.ID, DisplayName: e.DisplayName}
}
