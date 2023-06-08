package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"

	"github.com/rulanugrh/graphql/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name *string, email *string, username *string, password *string) (*model.User, error) {
	user := model.User{
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}

	return &user, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, id *int, paid *bool, total *int, userID *int) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, id *int, name *string, price *int, stock *int, author *string, examplar *int, sellerid *int, categoryid *int) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: CreateBook - createBook"))
}

// CreaeteCategory is the resolver for the creaeteCategory field.
func (r *mutationResolver) CreaeteCategory(ctx context.Context, id *int, name *string, description *string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: CreaeteCategory - creaeteCategory"))
}

// Finduser is the resolver for the finduser field.
func (r *queryResolver) Finduser(ctx context.Context, id *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Finduser - finduser"))
}

// Bookfindid is the resolver for the bookfindid field.
func (r *queryResolver) Bookfindid(ctx context.Context, id *string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: Bookfindid - bookfindid"))
}

// Categoryfindid is the resolver for the categoryfindid field.
func (r *queryResolver) Categoryfindid(ctx context.Context, id *string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: Categoryfindid - categoryfindid"))
}

// Findorder is the resolver for the findorder field.
func (r *queryResolver) Findorder(ctx context.Context, id *string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: Findorder - findorder"))
}

// Listnotpaid is the resolver for the listnotpaid field.
func (r *queryResolver) Listnotpaid(ctx context.Context, id *string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: Listnotpaid - listnotpaid"))
}

// Deletebookid is the resolver for the deletebookid field.
func (r *queryResolver) Deletebookid(ctx context.Context, id *string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: Deletebookid - deletebookid"))
}

// Findallorder is the resolver for the findallorder field.
func (r *queryResolver) Findallorder(ctx context.Context, id *string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: Findallorder - findallorder"))
}

// Listcart is the resolver for the listcart field.
func (r *queryResolver) Listcart(ctx context.Context, userID *string) (*model.OrderItem, error) {
	panic(fmt.Errorf("not implemented: Listcart - listcart"))
}

// Getallbook is the resolver for the getallbook field.
func (r *queryResolver) Getallbook(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented: Getallbook - getallbook"))
}

// Getallcategory is the resolver for the getallcategory field.
func (r *queryResolver) Getallcategory(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: Getallcategory - getallcategory"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
func (r *queryResolver) Order(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: Order - order"))
}
func (r *queryResolver) Book(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented: Book - book"))
}
