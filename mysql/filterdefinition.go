package mysql

// import "strings"

// type FilterDefinition struct {
// 	FieldName  string
// 	Values     []string
// 	FilterType FilterType
// 	Items      []FilterDefinition
// }

// type FilterType int

// const (
// 	And         FilterType = 1
// 	Or          FilterType = 2
// 	Equals      FilterType = 4
// 	NonEquals   FilterType = 8
// 	GreaterThan FilterType = 16
// 	Like        FilterType = 32
// )

// func (f FilterDefinition) GetFilterQuery() string {

// 	switch f.FilterType {
// 	case And:
// 		var s []string
// 		for _, item := range f.Items {
// 			s = append(s, item.GetFilterQuery())
// 		}

// 		return strings.Join(s, " AND ")
// 	case Or:
// 		var s []string
// 		for _, item := range f.Items {
// 			s = append(s, item.GetFilterQuery())
// 		}

// 		return strings.Join(s, " OR ")
// 	case Equals:
// 		if len(f.Values) == 1 {
// 			return f.FieldName + " = " + f.Values[0]
// 		} else {
// 			return f.FieldName + " IN (" + strings.Join(f.Values, ",") + ")"
// 		}
// 	case NonEquals:
// 		if len(f.Values) == 1 {
// 			return f.FieldName + " != " + f.Values[0]
// 		} else {
// 			return f.FieldName + " NOT IN (" + strings.Join(f.Values, ",") + ")"
// 		}
// 	case Like:
// 		if len(f.Values) == 1 {
// 			return f.FieldName + " LIKE " + f.Values[0]
// 		} else {
// 			s := ""
// 			for _, val := range f.Values {
// 				s += f.FieldName + " LIKE " + val
// 			}
// 			return s
// 		}
// 	case GreaterThan:
// 		if len(f.Values) == 1 {
// 			return f.FieldName + " > " + f.Values[0]
// 		} else {
// 			s := ""
// 			for _, val := range f.Values {
// 				s += f.FieldName + " > " + val
// 			}
// 			return s
// 		}
// 	}

// 	return ""
// }
