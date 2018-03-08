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
type productbacklog struct {
	id         graphqlgo.ID
	desc       string
	type_cd    string
	priority   string
	user_id    int32
	product_id int32
	storys     []*story
	user       *user
	product    *product
}

// Struct for upserting
type productbacklogInput struct {
	Id        *graphqlgo.ID
	Desc      string
	TypeCd    string
	Priority  string
	UserId    *int32
	ProductId *int32
	Storys    *[]storyInput
}

// Struct for response
type productbacklogResolver struct {
	productbacklog *productbacklog
}

func ResolveProductBackLog(args struct {
	ID graphqlgo.ID
}) (response []*productbacklogResolver) {
	if args.ID != "" {
		response = append(response, &productbacklogResolver{productbacklog: MapProductBackLog(models.GetProductBackLog(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllProductBackLogs(utils.GetDefaultLimitOffset()) {
		response = append(response, &productbacklogResolver{productbacklog: MapProductBackLog(val)})
	}
	return response
}

func ResolveCreateProductBackLog(args *struct {
	ProductBackLog *productbacklogInput
}) *productbacklogResolver {
	var productbacklog = &productbacklog{}

	if args.ProductBackLog.Id == nil {
		productbacklog = MapProductBackLog(models.PostProductBackLog(ReverseMapProductBackLog(args.ProductBackLog)))
	} else {
		productbacklog = MapProductBackLog(models.PutProductBackLog(ReverseMapProductBackLog(args.ProductBackLog)))
	}
	if productbacklog != nil && args.ProductBackLog.Storys != nil {
		for _, dev := range *args.ProductBackLog.Storys {
			if dev.Id == nil {
				story := ReverseMapStory(&dev)
				if story.ProductBackLogId != 0 && utils.ConvertId(productbacklog.id) != story.ProductBackLogId {
					// todo throw error
					return &productbacklogResolver{}
				}
				story.ProductBackLogId = utils.ConvertId(productbacklog.id)
				productbacklog.storys = append(productbacklog.storys, MapStory(models.PostStory(story)))
			} else {
				story := ReverseMapStory(&dev)
				if story.ProductBackLogId != 0 && utils.ConvertId(productbacklog.id) != story.ProductBackLogId {
					// todo throw error
					return &productbacklogResolver{}
				}
				story.ProductBackLogId = utils.ConvertId(productbacklog.id)
				productbacklog.storys = append(productbacklog.storys, MapStory(models.PutStory(story)))
			}
		}
	}
	return &productbacklogResolver{productbacklog}
}

// For Delete
func ResolveDeleteProductBackLog(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.ProductBackLogChildren) == 0 && len(models.ProductBackLogInterRelation) == 0 {
		del = models.DeleteProductBackLog(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.ProductBackLog
		database.SQL.Model(models.ProductBackLog{}).Preload("Storys").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		for _, v := range data.Storys {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteStory(args, "")
			count++
		}

		del = models.DeleteProductBackLog(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.ProductBackLog
	database.SQL.Model(models.ProductBackLog{}).Preload("Storys").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	for _, v := range data.Storys {
		if v.Id != 0 {
			flag++
		}
	}

	if flag == 0 {
		del = models.DeleteProductBackLog(utils.ConvertId(tempID), name)
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
func (r *productbacklogResolver) Id() graphqlgo.ID {
	return r.productbacklog.id
}
func (r *productbacklogResolver) Desc() string {
	return r.productbacklog.desc
}
func (r *productbacklogResolver) TypeCd() string {
	return r.productbacklog.type_cd
}
func (r *productbacklogResolver) Priority() string {
	return r.productbacklog.priority
}
func (r *productbacklogResolver) UserId() int32 {
	return r.productbacklog.user_id
}
func (r *productbacklogResolver) ProductId() int32 {
	return r.productbacklog.product_id
}
func (r *productbacklogResolver) Storys() []*storyResolver {
	var storys []*storyResolver
	if r.productbacklog != nil {
		story := models.GetStorysOfProductBackLog(utils.ConvertId(r.productbacklog.id))
		for _, value := range story {
			storys = append(storys, &storyResolver{MapStory(value)})
		}
		return storys
	}
	for _, value := range r.productbacklog.storys {
		storys = append(storys, &storyResolver{value})
	}
	return storys
}
func (r *productbacklogResolver) User() *userResolver {
	if r.productbacklog != nil {
		user := models.GetUserOfProductBackLog(ReverseMap2ProductBackLog(r.productbacklog))
		return &userResolver{MapUser(user)}
	}
	return &userResolver{r.productbacklog.user}
}
func (r *productbacklogResolver) Product() *productResolver {
	if r.productbacklog != nil {
		product := models.GetProductOfProductBackLog(ReverseMap2ProductBackLog(r.productbacklog))
		return &productResolver{MapProduct(product)}
	}
	return &productResolver{r.productbacklog.product}
}

// Mapper methods
func MapProductBackLog(modelProductBackLog models.ProductBackLog) *productbacklog {

	if reflect.DeepEqual(modelProductBackLog, models.ProductBackLog{}) {
		return &productbacklog{}
	}

	// Create graphql productbacklog from models ProductBackLog
	productbacklog := productbacklog{
		desc:       modelProductBackLog.Desc,
		id:         utils.UintToGraphId(modelProductBackLog.Id),
		priority:   modelProductBackLog.Priority,
		product_id: int32(modelProductBackLog.ProductId),
		type_cd:    modelProductBackLog.TypeCd,
		user_id:    int32(modelProductBackLog.UserId),
	}
	return &productbacklog
}

// Reverse Mapper methods
func ReverseMapProductBackLog(mygraphqlProductBackLog *productbacklogInput) models.ProductBackLog {

	if reflect.DeepEqual(mygraphqlProductBackLog, productbacklogInput{}) {
		return models.ProductBackLog{}
	}

	// Create graphql productbacklog from models ProductBackLog
	var productbacklogModel models.ProductBackLog
	if mygraphqlProductBackLog.Id == nil {
		productbacklogModel = models.ProductBackLog{
			Desc:     mygraphqlProductBackLog.Desc,
			Priority: mygraphqlProductBackLog.Priority,
			TypeCd:   mygraphqlProductBackLog.TypeCd,
		}
	} else {
		productbacklogModel = models.ProductBackLog{
			Desc:      mygraphqlProductBackLog.Desc,
			Id:        utils.ConvertId(*mygraphqlProductBackLog.Id),
			Priority:  mygraphqlProductBackLog.Priority,
			ProductId: utils.Int32ToUint(*mygraphqlProductBackLog.ProductId),
			TypeCd:    mygraphqlProductBackLog.TypeCd,
			UserId:    utils.Int32ToUint(*mygraphqlProductBackLog.UserId),
		}
	}
	return productbacklogModel
}
func ReverseMap2ProductBackLog(structProductBackLog *productbacklog) models.ProductBackLog {

	if reflect.DeepEqual(structProductBackLog, productbacklog{}) {
		return models.ProductBackLog{}
	}

	// Create graphql productbacklog from models ProductBackLog
	modelProductBackLog := models.ProductBackLog{
		Desc:      structProductBackLog.desc,
		Id:        utils.ConvertId(structProductBackLog.id),
		Priority:  structProductBackLog.priority,
		ProductId: uint(structProductBackLog.product_id),
		TypeCd:    structProductBackLog.type_cd,
		UserId:    uint(structProductBackLog.user_id),
	}
	return modelProductBackLog
}
