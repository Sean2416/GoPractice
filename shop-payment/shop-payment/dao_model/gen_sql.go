package dao_model

import (
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"shop-payment/shop-payment/sutil"
	"reflect"
	"sort"
	"strings"
)

const (
	whereTag  = "where"
	ignoreTag = "-"
	emptyTag  = ""
)

func GenSql(table string, columns []string, condition interface{}, tags []string, page *Page, orders map[string]string, groups []string) (sql string, args []interface{}, err error) {
	if sutil.IsEmptyValue(reflect.ValueOf(table)) {
		err = errors.New("empty table name")
		return
	}

	return buildSelectBy(table, columns, condition, tags, page, orders, groups)
}

// 建構select 語法
func buildSelectBy(table string, columns []string, condition interface{}, tags []string, page *Page, order map[string]string, groups []string) (sql string, args []interface{}, err error) {

	sb := sqlbuilder.NewSelectBuilder().From(table)

	// 構造 columns
	if !sutil.IsEmptyValue(reflect.ValueOf(columns)) {
		sb.Select(columns...)
	} else {
		sb.Select("*")
	}

	// 查詢邏輯
	if condition != nil && !sutil.IsEmptyValue(reflect.ValueOf(condition)) {
		var params map[string]interface{}

		if len(tags) == 0 {
			params = mapFromTag(condition, whereTag)
		} else {
			params = mapFromTags(condition, tags...)
		}
		if len(params) > 0 {
			var wheres []string
			wheres, err = buildWhere(params, &sb.Cond)
			if err != nil {
				return
			}
			sb.Where(wheres...)
		}
	}

	// 排序
	if !sutil.IsEmptyValue(reflect.ValueOf(order)) {
		var columns []string
		for column, dir := range order {
			columns = append(columns, column+" "+strings.ToUpper(dir))
		}
		sort.Strings(columns)
		sb.OrderBy(columns...)
	}

	// 分頁
	if page == nil || page.Limit <= 0 {
		sb.Limit(-1).Offset(-1)
	} else {
		sb.Limit(page.Limit).Offset(page.Offset)
	}

	if len(groups) > 0 {
		sb.GroupBy(groups...)
	}

	sql, args = sb.Build()
	return
}

func derefType(t reflect.Type) reflect.Type {
	for k := t.Kind(); k == reflect.Ptr || k == reflect.Interface; k = t.Kind() {
		t = t.Elem()
	}
	return t
}

func derefValue(value interface{}) reflect.Value {
	v := reflect.ValueOf(value)

	for k := v.Kind(); k == reflect.Ptr || k == reflect.Interface; k = v.Kind() {
		v = v.Elem()
	}

	return v
}

// 取得struct tag標示的k-v map
// 兼容map类型
func mapFromTag(obj interface{}, tagName string) map[string]interface{} {
	aliases := make(map[string]interface{})
	if reflect.ValueOf(obj).Kind() == reflect.Map {
		aliases, ok := obj.(map[string]interface{})
		if ok {
			return aliases
		} else {
			return nil
		}
	}
	v := derefValue(obj)
	t := derefType(reflect.TypeOf(obj))
	len := t.NumField()

	for i := 0; i < len; i++ {
		field := t.Field(i)

		// 忽略匿名變量
		if field.Anonymous {
			continue
		}

		// 忽略 ptr = nil的成員
		fieldValue := v.FieldByName(field.Name)
		//if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
		if sutil.IsEmptyValue(fieldValue) {
			continue
		}

		data := derefValue(fieldValue.Interface()).Interface()

		tagValue := field.Tag.Get(tagName)
		if tagValue == ignoreTag || tagValue == emptyTag {
			continue
		}

		tags := strings.Split(tagValue, ",")
		for _, tag := range tags {
			if tag != "" {
				aliases[tag] = data
			}
		}

	}
	return aliases
}

var opMap = map[string]string{
	"=":        "Equal",
	"!=":       "NotEqual",
	">":        "GreaterThan",
	"<":        "LessThan",
	"<=":       "LessEqualThan",
	">=":       "GreaterEqualThan",
	"like":     "Like",
	"not like": "NotLike",
	"in":       "In",
	"not in":   "NotIn",
}

var operators = []string{
	"not like", "not in", "like", "in",
	">=", "<=", "!=", ">", "<", "=",
}

func mapFromTags(structValue interface{}, tags ...string) map[string]interface{} {
	allFields := make(map[string]interface{})
	for i := 0; i < len(tags); i++ {
		fields := mapFromTag(structValue, tags[i])
		for k, v := range fields {
			allFields[k] = v
		}
	}
	return allFields
}

func ToSlice(input interface{}) (slice []interface{}, err error) {
	value := reflect.ValueOf(input)
	if value.Kind() != reflect.Slice {
		err = errors.New("input should be slice")
	}
	len := value.Len()
	slice = make([]interface{}, len)
	for i := 0; i < len; i++ {
		slice[i] = value.Index(i).Interface()
	}
	return
}

// 根據名稱動態call obj方法. 功能同PHP call_user_func_array()
func invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

// where 條件表達式解析
func parseExpr(expr string) (field string, operator string, err error) {
	if len(expr) == 0 {

		err = errors.New("empty expr")
		return
	}
	for _, op := range operators {
		expr = strings.TrimSpace(expr)
		if strings.HasSuffix(expr, op) {
			field = strings.TrimSpace(strings.TrimSuffix(expr, op))
			operator = strings.ToLower(strings.TrimSpace(op))
			return
			//break
		}
	}
	err = fmt.Errorf("invalid expression: %v", expr)
	return
}

// where条件. 依賴sqlbuilder.Cond
func buildWhere(params map[string]interface{}, cond *sqlbuilder.Cond) (wheres []string, err error) {
	var column, operator string
	for key, value := range params {
		column, operator, err = parseExpr(key)
		if err != nil {
			continue
		}

		// in & not in 參數為array, 額外處理
		if strings.HasSuffix(operator, "in") {
			args := []interface{}{}
			vs := reflect.ValueOf(value)
			if vs.Kind() == reflect.Slice || vs.Kind() == reflect.Array {
				var where string
				args, err = ToSlice(value)
				if strings.Index(operator, "not") == -1 {
					where = cond.In(column, args...)
				} else {
					where = cond.NotIn(column, args...)
				}
				wheres = append(wheres, where)
			} else {

				err = errors.New("invalid `in` argument type")
				return
			}

		} else {
			wheres = append(wheres, invoke(cond, opMap[operator], column, value)[0].String())
		}
	}
	// 條件有序
	if len(wheres) > 0 {
		sort.Strings(wheres)
	}
	return
}
