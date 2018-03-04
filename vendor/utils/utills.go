package utils

import (
	"strconv"
	"github.com/neelance/graphql-go"
	"strings"
	"log"
	"fmt"
)

func StringToUInt(ID string) uint {

	if ID == ""{
		return 0
	}else{

		u64, err := strconv.ParseUint(ID, 10, 32)
		if err != nil {
			log.Println(err)
			return 0
		}
		wd := uint(u64)
		return wd
	}

}

func ConvertId(id graphql.ID) uint {

	if id == ""{
		return 0
	}else{

		val := StringToUInt(string(id))
		return val
	}

}

func UintToGraphId(ID uint) graphql.ID {

	if ID == 0{
		return ""
	}else{
		str := fmt.Sprint(ID)
		return graphql.ID(str)

	}

}
func RuneToGraphId(ID rune) graphql.ID {
	str := fmt.Sprint(ID)
	return graphql.ID(str)
}

func Int32ToUint(ID int32) uint {

	if ID == 0{
		return 0
	}else{

		str := fmt.Sprint(ID)
		str2 := StringToUInt(str)
		return str2
	}
}

func SAppend(old *string, new string) {
	*old = fmt.Sprintf("%s %s", *old, new)
}

func IsValueInList(value string, list []string) bool {
	for _, v := range list {
		if strings.ToLower(v) == strings.ToLower(value) {
			return true
		}
	}
	return false
}
