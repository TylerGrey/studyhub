package repo

import (
	"fmt"
	"reflect"
)

// Order ...
type Order struct {
	Field     string
	Direction string
}

// ListArgs ...
type ListArgs struct {
	First  *int32
	Last   *int32
	After  *string
	Before *string
	Order  *Order
}

// PageInfo ...
type PageInfo struct {
	HasNext bool
	HasPrev bool
	Total   int32
}

func getBaseCursor(args ListArgs) string {
	var cursor = defaultListCursor

	if args.Order != nil {
		var cursorField string
		if args.Order.Field == "CREATED_AT" {
			cursorField = "UNIX_TIMESTAMP(created_at)"
		} else {
			cursorField = args.Order.Field
		}

		if args.Order.Direction == "DESC" {
			cursor = fmt.Sprintf("CONCAT(%s, LPAD(id, 10, '0'))", cursorField)
		} else {
			cursor = fmt.Sprintf("CONCAT(POW(10, 10) - %s, LPAD(POW(10, 10) - id, 10, '0'))", cursorField)
		}
	}

	return cursor
}

func getOrderBy(args ListArgs) (string, string) {
	var (
		field     = defaultListOrderField
		direction = defaultListOrderDirection
	)

	if args.Order != nil {
		field = args.Order.Field
		direction = args.Order.Direction
		if args.Last != nil {
			if direction == "DESC" {
				direction = "ASC"
			} else {
				direction = "DESC"
			}
		}
	}

	return field, direction
}

func getCursor(args ListArgs) (string, string) {
	var (
		cursor    string
		direction = "<"
	)

	if (args.First != nil && args.Before != nil) || (args.Last != nil && args.After != nil) {
		direction = ">"
	}

	if args.After != nil {
		cursor = *args.After
	} else if args.Before != nil {
		cursor = *args.Before
	}

	return cursor, direction
}

func getLimit(args ListArgs) int32 {
	if args.First != nil {
		return *args.First
	} else if args.Last != nil {
		return *args.Last
	}

	return defaultListLimit
}

func getPageInfo(args ListArgs, limit int32, list interface{}) (interface{}, PageInfo) {
	var (
		hasNext = false
		hasPrev = false
	)

	elems := reflect.ValueOf(list).Elem()
	hasMore := elems.Len() > int(limit)
	if hasMore {
		if args.Before != nil {
			hasPrev = true
			elems.Set(elems.Slice(1, elems.Len()))
		} else {
			hasNext = true
			elems.Set(elems.Slice(0, elems.Len()-1))
		}
	}

	if args.Before != nil {
		hasNext = true
	} else if args.After != nil {
		hasPrev = true
	}

	return elems.Interface(), PageInfo{
		HasNext: hasNext,
		HasPrev: hasPrev,
	}
}
