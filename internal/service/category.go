package service

import (
	"context"
	"io"

	"github.com.br/leomoritz/gRPC/internal/database"
	"github.com.br/leomoritz/gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

// Construtor do serviço com instância do banco para category
func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})

	}
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) UpdateCategory(ctx context.Context, in *pb.UpdateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Find(in.Category.Id)
	if err != nil {
		return nil, err
	}

	category.Name = in.Category.Name
	category.Description = in.Category.Description

	categoryUpdated, err := c.CategoryDB.Update(category)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          categoryUpdated.ID,
		Name:        categoryUpdated.Name,
		Description: categoryUpdated.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.DeleteCategoryResponse, error) {
	response, err := c.CategoryDB.Delete(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteCategoryResponse{
		IsDeleted: response,
	}, nil
}
