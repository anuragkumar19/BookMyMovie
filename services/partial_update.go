package services

// import (
// 	"slices"
// )

// type UpdateField[T any] struct {
// 	Value T
// 	valid bool
// }

// func (f UpdateField[T]) Valid() bool {
// 	return f.valid
// }

// func (f UpdateField[T]) Update(value T) T {
// 	if f.valid {
// 		return f.Value
// 	}
// 	return value
// }

// func NewUpdateField[T any](value T, valid bool) UpdateField[T] {
// 	return UpdateField[T]{
// 		Value: value,
// 		valid: valid,
// 	}
// }

// func NewValidUpdateField[T any](value T) UpdateField[T] {
// 	return UpdateField[T]{
// 		Value: value,
// 		valid: true,
// 	}
// }

// type Action string

// const (
// 	ActionIgnore  Action = ""
// 	ActionReplace Action = "replace"
// 	ActionAppend  Action = "append"
// 	ActionRemove  Action = "remove"
// )

// type UpdateFieldWithAction[T comparable] struct {
// 	Value  []T
// 	action Action
// }

// func NewUpdateFieldWithAction[T comparable](value []T, action Action) UpdateFieldWithAction[T] {
// 	switch action {
// 	case ActionReplace, ActionAppend, ActionRemove:
// 	default:
// 		action = ActionIgnore
// 	}
// 	return UpdateFieldWithAction[T]{
// 		Value:  value,
// 		action: action,
// 	}
// }

// func (f UpdateFieldWithAction[T]) Update(value []T) []T {
// 	switch f.action {
// 	case ActionAppend:
// 		{
// 			newS := make([]T, 0, len(value)+len(f.Value))
// 			newS = append(newS, value...)
// 			newS = append(newS, f.Value...)
// 			return newS
// 		}
// 	case ActionRemove:
// 		{
// 			fields := make(map[T]bool)
// 			for _, v := range f.Value {
// 				fields[v] = true
// 			}

// 			c := slices.Clone(value)
// 			return slices.DeleteFunc(c, func(v T) bool {
// 				return fields[v]
// 			})
// 		}

// 	case ActionReplace:
// 		return f.Value
// 	default:
// 		return value
// 	}
// }

// func (f UpdateFieldWithAction[T]) Valid() bool {
// 	switch f.action {
// 	case ActionAppend, ActionReplace, ActionRemove:
// 		return true
// 	default:
// 		return false
// 	}
// }
