package unit_test

import (
	"app/api/models"
	"context"
	"strconv"
	"testing"
)

func TestCreateBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateBrand
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateBrand{
				Brand_name: "Test Brand",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := brandTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.BrandPrimaryKey
		Output  *models.Brand
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.BrandPrimaryKey{
				Brand_id: 10,
			},
			Output: &models.Brand{
				Brand_id:   10,
				Brand_name: "string",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			brand, err := brandTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if brand.Brand_name != test.Output.Brand_name || brand.Brand_id != test.Output.Brand_id {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *brand, *test.Output)
				return
			}

		})
	}
}

func TestUpdateBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateBrand
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateBrand{
				Brand_id:   10,
				Brand_name: "Test Brand updated",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := brandTestRepo.Update(context.Background(), test.Input)

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

func TestDeleteBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.BrandPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.BrandPrimaryKey{
				Brand_id: 10,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := brandTestRepo.Delete(context.Background(), test.Input)

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
