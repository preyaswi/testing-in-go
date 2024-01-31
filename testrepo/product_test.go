package test

import (
	"ks/domain"
	"ks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestAddProduct(t *testing.T)  {
	testProduct:=domain.Product{
		ProductName: "masccara",
		Price: 234,
		Catergory: "makeup",
	}
	err:=repository.GetProductByName(testProduct)
	assert.NoError(t,err,"expected no error")

	err=repository.CreateProduct(testProduct)
	assert.NoError(t,err,"expected no error")

	err=repository.DeleteProduct(testProduct)
	assert.NoError(t,err,"expected no error")
}
func TestGetAllProduct(t *testing.T)  {
	expectedProduct:=[]domain.Product{
		{ProductName: "baniyan",Price: 56,Catergory: "men"},
	}
	list,err:=repository.GetAllProduct()
	assert.NoError(t,err,"expected no erorr")
	assert.Equal(t,expectedProduct,list,"returned products do no match expected products")

}