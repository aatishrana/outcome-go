package mygraphql

import (
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
}

// Struct for upserting
type productbacklogInput struct {
	Id        *graphqlgo.ID
	Desc      string
	TypeCd    string
	Priority  string
	UserId    *int32
	ProductId *int32
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
	for _, val := range models.GetAllProductBackLogs() {
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
