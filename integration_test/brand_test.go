package integration_test

import (
	"app/api/models"
	"fmt"
	"strconv"
	"sync"

	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestBrand(t *testing.T) {
	s := 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createBrand(t)
			updateBrand(t, id)
			deleteBrand(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createBrand(t *testing.T) int {
	response := &models.Brand{}

	request := &models.CreateBrand{
		Brand_name: faker.Name(),
	}
	resp, err := PerformRequest(http.MethodPost, "/brand", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Brand_id
}

func updateBrand(t *testing.T, id int) int {
	response := &models.Brand{}

	request := &models.UpdateBrand{
		Brand_name: faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPut, "/category/"+strconv.Itoa(id), request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 202)
	}

	return response.Brand_id
}

func deleteBrand(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, "/brand/"+strconv.Itoa(id), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 202)
	}

	return 0
}
