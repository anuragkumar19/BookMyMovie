package services

import (
	"slices"
)

type UpdateField[T any] struct {
	value T
	valid bool
}

func (f UpdateField[T]) Value(value T) T {
	if f.valid {
		return f.value
	}
	return value
}

func NewUpdateField[T any](value T, valid bool) UpdateField[T] {
	return UpdateField[T]{
		value: value,
		valid: valid,
	}
}

func NewValidUpdateField[T any](value T) UpdateField[T] {
	return UpdateField[T]{
		value: value,
		valid: true,
	}
}

type Action string

const (
	ActionIgnore  Action = ""
	ActionReplace Action = "replace"
	ActionAppend  Action = "append"
	ActionRemove  Action = "remove"
)

type UpdateFieldWithAction[T comparable] struct {
	value  []T
	action Action
}

func NewUpdateFieldWithAction[T comparable](value []T, action Action) UpdateFieldWithAction[T] {
	switch action {
	case ActionReplace, ActionAppend, ActionRemove:
	default:
		action = ActionIgnore
	}
	return UpdateFieldWithAction[T]{
		value:  value,
		action: action,
	}
}

func (f UpdateFieldWithAction[T]) Value(value []T) []T {
	switch f.action {
	case ActionAppend:
		{
			newS := make([]T, 0, len(value)+len(f.value))
			newS = append(newS, value...)
			newS = append(newS, f.value...)
			return newS
		}
	case ActionRemove:
		{
			fields := make(map[T]bool)
			for _, v := range f.value {
				fields[v] = true
			}

			c := slices.Clone(value)
			return slices.DeleteFunc(c, func(v T) bool {
				return fields[v]
			})
		}

	case ActionReplace:
		return f.value
	default:
		return value
	}
}
