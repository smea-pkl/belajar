package main

import "context"

// 1. Interface

type User struct {
	Name     string
	Username string
}

type FilterData interface {
	Filter(user User) (bool, error)
}

type FilterBaru interface {
	FilterV2(ctx context.Context, user User) (bool, error)
}

type filterReyhan struct{}

// Filter implements [FilterData].
func (f filterReyhan) Filter(user User) (bool, error) {
	// 1. filter ruwet
	return true, nil
}

// FilterV2 implements [FilterBaru].
func (f filterReyhan) FilterV2(ctx context.Context, user User) (bool, error) {
	panic("unimplemented")
}

type filterHeri struct{}

// FilterV2 implements [FilterBaru].
func (f filterHeri) FilterV2(ctx context.Context, user User) (bool, error) {
	panic("unimplemented")
}

// Filter implements [FilterData].
func (f filterHeri) Filter(user User) (bool, error) {
	// 1. filter ruwet
	return true, nil
}

type GetUser func(userid int) (User, error)

func main() {

}

func getUserHeriVersion(userid int) (User, error) {
	return User{}, nil
}
