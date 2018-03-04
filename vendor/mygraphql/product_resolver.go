package mygraphql

import (
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type product struct {
	id      graphqlgo.ID
	name    string
	desc    string
	user_id int32
	org_id  int32
	org     *org
	user    *user
}

// Struct for upserting
type productInput struct {
	Id     *graphqlgo.ID
	Name   string
	Desc   string
	UserId *int32
	OrgId  *int32
}

// Struct for response
type productResolver struct {
	product *product
}

func ResolveProduct(args struct {
	ID graphqlgo.ID
}) (response []*productResolver) {
	if args.ID != "" {
		response = append(response, &productResolver{product: MapProduct(models.GetProduct(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllProducts() {
		response = append(response, &productResolver{product: MapProduct(val)})
	}
	return response
}

func ResolveCreateProduct(args *struct {
	Product *productInput
}) *productResolver {
	var product = &product{}

	if args.Product.Id == nil {
		product = MapProduct(models.PostProduct(ReverseMapProduct(args.Product)))
	} else {
		product = MapProduct(models.PutProduct(ReverseMapProduct(args.Product)))
	}
	return &productResolver{product}
}

// For Delete
func ResolveDeleteProduct(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.ProductChildren) == 0 && len(models.ProductInterRelation) == 0 {
		del = models.DeleteProduct(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	return response
}

// Fields resolvers
func (r *productResolver) Id() graphqlgo.ID {
	return r.product.id
}
func (r *productResolver) Name() string {
	return r.product.name
}
func (r *productResolver) Desc() string {
	return r.product.desc
}
func (r *productResolver) UserId() int32 {
	return r.product.user_id
}
func (r *productResolver) OrgId() int32 {
	return r.product.org_id
}
func (r *productResolver) Org() *orgResolver {
	if r.product != nil {
		org := models.GetOrgOfProduct(ReverseMap2Product(r.product))
		return &orgResolver{MapOrg(org)}
	}
	return &orgResolver{r.product.org}
}
func (r *productResolver) User() *userResolver {
	if r.product != nil {
		user := models.GetUserOfProduct(ReverseMap2Product(r.product))
		return &userResolver{MapUser(user)}
	}
	return &userResolver{r.product.user}
}

// Mapper methods
func MapProduct(modelProduct models.Product) *product {

	if reflect.DeepEqual(modelProduct, models.Product{}) {
		return &product{}
	}

	// Create graphql product from models Product
	product := product{
		desc:    modelProduct.Desc,
		id:      utils.UintToGraphId(modelProduct.Id),
		name:    modelProduct.Name,
		org_id:  int32(modelProduct.OrgId),
		user_id: int32(modelProduct.UserId),
	}
	return &product
}

// Reverse Mapper methods
func ReverseMapProduct(mygraphqlProduct *productInput) models.Product {

	if reflect.DeepEqual(mygraphqlProduct, productInput{}) {
		return models.Product{}
	}

	// Create graphql product from models Product
	var productModel models.Product
	if mygraphqlProduct.Id == nil {
		productModel = models.Product{
			Desc: mygraphqlProduct.Desc,
			Name: mygraphqlProduct.Name,
		}
	} else {
		productModel = models.Product{
			Desc:   mygraphqlProduct.Desc,
			Id:     utils.ConvertId(*mygraphqlProduct.Id),
			Name:   mygraphqlProduct.Name,
			OrgId:  utils.Int32ToUint(*mygraphqlProduct.OrgId),
			UserId: utils.Int32ToUint(*mygraphqlProduct.UserId),
		}
	}
	return productModel
}
func ReverseMap2Product(structProduct *product) models.Product {

	if reflect.DeepEqual(structProduct, product{}) {
		return models.Product{}
	}

	// Create graphql product from models Product
	modelProduct := models.Product{
		Desc:   structProduct.desc,
		Id:     utils.ConvertId(structProduct.id),
		Name:   structProduct.name,
		OrgId:  uint(structProduct.org_id),
		UserId: uint(structProduct.user_id),
	}
	return modelProduct
}
