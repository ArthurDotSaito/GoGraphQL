// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Query struct {
}

type NewCategory struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type NewCourse struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	CategoryID  string  `json:"categoryId"`
}
