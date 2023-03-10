package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/dgmo/gqlgo/server/core"
	"github.com/dgmo/gqlgo/server/graph/models"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input models.CreateProductInput) (models.CreateProductResult, error) {
	if product, err := core.CreateProduct(ctx, input); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input models.UpdateProductInput) (models.UpdateProductResult, error) {
	if product, err := core.UpdateProduct(ctx, id, input); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (models.DeleteProductResult, error) {
	if product, err := core.DeleteProduct(ctx, id); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input models.CreateCategoryInput) (models.CreateCategoryResult, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, input models.UpdateCategoryInput) (models.UpdateCategoryResult, error) {
	panic(fmt.Errorf("not implemented: UpdateCategory - updateCategory"))
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (models.DeleteCategoryResult, error) {
	panic(fmt.Errorf("not implemented: DeleteCategory - deleteCategory"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, first *int, after *string, last *int, before *string, filter *models.ProductFilter, sorting *models.ProductSort) (*models.ProductConnection, error) {
	if products, err := core.Products(ctx, first, after, last, before, filter, sorting); err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*models.Product, error) {
	product, err := core.Product(ctx, id)
	if err != nil {
		return nil, err
	} else {
		return product, nil
	}
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context, first *int, after *string, last *int, before *string, filter *models.CategoryFilter, sorting *models.CategorySort) (*models.CategoryConnection, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*models.Category, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// ProductCreated is the resolver for the productCreated field.
func (r *subscriptionResolver) ProductCreated(ctx context.Context) (<-chan *models.Product, error) {
	panic(fmt.Errorf("not implemented: ProductCreated - productCreated"))
}

// ProductUpdated is the resolver for the productUpdated field.
func (r *subscriptionResolver) ProductUpdated(ctx context.Context) (<-chan *models.Product, error) {
	panic(fmt.Errorf("not implemented: ProductUpdated - productUpdated"))
}

// ProductDeleted is the resolver for the productDeleted field.
func (r *subscriptionResolver) ProductDeleted(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: ProductDeleted - productDeleted"))
}

// CategoryCreated is the resolver for the categoryCreated field.
func (r *subscriptionResolver) CategoryCreated(ctx context.Context) (<-chan *models.Category, error) {
	panic(fmt.Errorf("not implemented: CategoryCreated - categoryCreated"))
}

// CategoryUpdated is the resolver for the categoryUpdated field.
func (r *subscriptionResolver) CategoryUpdated(ctx context.Context) (<-chan *models.Category, error) {
	panic(fmt.Errorf("not implemented: CategoryUpdated - categoryUpdated"))
}

// CategoryDeleted is the resolver for the categoryDeleted field.
func (r *subscriptionResolver) CategoryDeleted(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: CategoryDeleted - categoryDeleted"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
