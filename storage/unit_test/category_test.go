package unit_test

import (
	"app/api/models"
	"context"
	"strconv"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateCategory
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateCategory{
				Category_name: "Test Category",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := categoryTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if i, _ := strconv.Atoi(id); i <= 0 {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestGetByIdCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CategoryPrimaryKey
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CategoryPrimaryKey{
				Category_id: 9,
			},
			Output: &models.Category{
				Category_id:   9,
				Category_name: "Test Category",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			category, err := categoryTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if category.Category_name != test.Output.Category_name || category.Category_id != test.Output.Category_id {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *category, *test.Output)
				return
			}

		})
	}
}

func TestUpdateCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateCategory
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateCategory{
				Category_id:   10,
				Category_name: "Test Category updated",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := categoryTestRepo.Update(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}

func TestDeleteCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CategoryPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CategoryPrimaryKey{
				Category_id: 10,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := categoryTestRepo.Delete(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}
