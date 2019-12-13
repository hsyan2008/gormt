package genmysql

import (
	"strings"

	"github.com/hsyan2008/gormt/data/view/model"
)

// filterModel filter.过滤 gorm.Model
func filterModel(list *[]genColumns) bool {
	var _temp []genColumns

	num := 0
	for _, v := range *list {
		if strings.EqualFold(v.Field, "id") ||
			strings.EqualFold(v.Field, "created_at") ||
			strings.EqualFold(v.Field, "updated_at") ||
			strings.EqualFold(v.Field, "deleted_at") {
			num++
		} else {
			_temp = append(_temp, v)
		}
	}

	if num >= 4 {
		*list = _temp
		return true
	}

	return false
}

// fixForeignKey fix foreign key.过滤外键
func fixForeignKey(list []genForeignKey, columuName string, result *[]model.ForeignKey) {
	for _, v := range list {
		if strings.EqualFold(v.ColumnName, columuName) { // find it .找到了
			*result = append(*result, model.ForeignKey{
				TableName:  v.ReferencedTableName,
				ColumnName: v.ReferencedColumnName,
			})
		}
	}
}
