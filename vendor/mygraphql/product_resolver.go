package mygraphql

import (
	database "database"
	fmt "fmt"
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type product struct {
	id              graphqlgo.ID
	name            string
	desc            string
	user_id         int32
	org_id          int32
	productbacklogs []*productbacklog
	projects        []*project
	org             *org
	user            *user
}

// Struct for upserting
type productInput struct {
	Id              *graphqlgo.ID
	Name            string
	Desc            string
	UserId          *int32
	OrgId           *int32
	ProductBackLogs *[]productbacklogInput
	Projects        *[]projectInput
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
	if product != nil && args.Product.ProductBackLogs != nil {
		for _, dev := range *args.Product.ProductBackLogs {
			if dev.Id == nil {
				productbacklog := ReverseMapProductBackLog(&dev)
				if productbacklog.ProductId != 0 && utils.ConvertId(product.id) != productbacklog.ProductId {
					// todo throw error
					return &productResolver{}
				}
				productbacklog.ProductId = utils.ConvertId(product.id)
				product.productbacklogs = append(product.productbacklogs, MapProductBackLog(models.PostProductBackLog(productbacklog)))
			} else {
				productbacklog := ReverseMapProductBackLog(&dev)
				if productbacklog.ProductId != 0 && utils.ConvertId(product.id) != productbacklog.ProductId {
					// todo throw error
					return &productResolver{}
				}
				productbacklog.ProductId = utils.ConvertId(product.id)
				product.productbacklogs = append(product.productbacklogs, MapProductBackLog(models.PutProductBackLog(productbacklog)))
			}
		}
	}
	if product != nil && args.Product.Projects != nil {
		for _, dev := range *args.Product.Projects {
			if dev.Id == nil {
				project := ReverseMapProject(&dev)
				if project.ProductId != 0 && utils.ConvertId(product.id) != project.ProductId {
					// todo throw error
					return &productResolver{}
				}
				project.ProductId = utils.ConvertId(product.id)
				product.projects = append(product.projects, MapProject(models.PostProject(project)))
			} else {
				project := ReverseMapProject(&dev)
				if project.ProductId != 0 && utils.ConvertId(product.id) != project.ProductId {
					// todo throw error
					return &productResolver{}
				}
				project.ProductId = utils.ConvertId(product.id)
				product.projects = append(product.projects, MapProject(models.PutProject(project)))
			}
		}
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
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.Product
		database.SQL.Model(models.Product{}).Preload("ProductBackLogs").Preload("Projects").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		for _, v := range data.ProductBackLogs {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteProductBackLog(args, "")
			count++
		}

		for _, v := range data.Projects {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteProject(args, "")
			count++
		}

		del = models.DeleteProduct(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.Product
	database.SQL.Model(models.Product{}).Preload("ProductBackLogs").Preload("Projects").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	for _, v := range data.ProductBackLogs {
		if v.Id != 0 {
			flag++
		}
	}

	for _, v := range data.Projects {
		if v.Id != 0 {
			flag++
		}
	}

	if flag == 0 {
		del = models.DeleteProduct(utils.ConvertId(tempID), name)
		count++
		response = &count
	} else {
		// show error
		fmt.Println("Cannot Delete :", tempID)
		del = false
		response = &count
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
func (r *productResolver) ProductBackLogs() []*productbacklogResolver {
	var productbacklogs []*productbacklogResolver
	if r.product != nil {
		productbacklog := models.GetProductBackLogsOfProduct(utils.ConvertId(r.product.id))
		for _, value := range productbacklog {
			productbacklogs = append(productbacklogs, &productbacklogResolver{MapProductBackLog(value)})
		}
		return productbacklogs
	}
	for _, value := range r.product.productbacklogs {
		productbacklogs = append(productbacklogs, &productbacklogResolver{value})
	}
	return productbacklogs
}
func (r *productResolver) Projects() []*projectResolver {
	var projects []*projectResolver
	if r.product != nil {
		project := models.GetProjectsOfProduct(utils.ConvertId(r.product.id))
		for _, value := range project {
			projects = append(projects, &projectResolver{MapProject(value)})
		}
		return projects
	}
	for _, value := range r.product.projects {
		projects = append(projects, &projectResolver{value})
	}
	return projects
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
