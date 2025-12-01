package core_test

// FIXME `stretchr/testify/assert` から脱却
import (
	"errors"
	"github.com/motojouya/mvc_gp/domain/basic/core"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestFilter(t *testing.T) {
	var list = []string{"this", "test", "item"}
	var predicate = func(item string) bool {
		var chars = []rune(item)
		if len(chars) > 0 && chars[0] == 't' {
			return true
		}
		return false
	}

	var tList = core.Filter(list, predicate)

	assert.Equal(t, 2, len(tList))
	assert.Equal(t, "this", tList[0])
	assert.Equal(t, "test", tList[1])

	t.Logf("filtered list: %v", tList)
}

func TestMap(t *testing.T) {
	var list = []string{"this", "test", "item"}
	var mapper = func(item string) string {
		return item + "_mapped"
	}

	var tList = core.Map(list, mapper)

	assert.Equal(t, 3, len(tList))
	assert.Equal(t, "this_mapped", tList[0])
	assert.Equal(t, "test_mapped", tList[1])
	assert.Equal(t, "item_mapped", tList[2])

	t.Logf("mapped list: %v", tList)
}

func TestFold(t *testing.T) {
	var list = []string{"this", "test", "item"}
	var folder = func(accumulator string, item string) (string, error) {
		return accumulator + "_" + item, nil
	}

	var result, err = core.Fold(list, "first", folder)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.Equal(t, "first_this_test_item", result)

	t.Logf("reduced result: %s", result)
}

func TestFoldError(t *testing.T) {
	var list = []string{"this", "test", "item"}
	var folder = func(accumulator string, item string) (string, error) {
		return "", errors.New("test error")
	}

	var _, err = core.Fold(list, "first", folder)
	if err == nil {
		t.Fatal("expected error, but got nil")
	}
}

func TestReduce(t *testing.T) {
	var list = []int{1, 2, 3}
	var reducer = func(accumulator int, item int) int {
		return accumulator + item
	}

	var result = core.Reduce(list, reducer)

	assert.Equal(t, 6, result)

	t.Logf("reduced result: %d", result)
}

func TestSome(t *testing.T) {
	var list1 = []string{"this", "test", "item"}
	var list2 = []string{"this", "tast", "item"}
	var predicate = func(item string) bool {
		return item == "test"
	}

	var contains1 = core.Some(list1, predicate)
	var contains2 = core.Some(list2, predicate)

	assert.True(t, contains1)
	assert.False(t, contains2)
}

func TestEvery(t *testing.T) {
	var list1 = []string{"this", "test", "item"}
	var list2 = []string{"this", "test", "temi"}
	var predicate = func(item string) bool {
		var chars = []rune(item)
		return chars[0] == 't'
	}

	var unMatched = core.Every(list1, predicate)
	var allMatched = core.Every(list2, predicate)

	assert.False(t, unMatched)
	assert.True(t, allMatched)
}

func TestFind(t *testing.T) {
	var list = []string{"this", "test", "item"}
	var predicate = func(item string) bool {
		var chars = []rune(item)
		return chars[0] == 't'
	}

	var foundItem, exists = core.Find(list, predicate)

	assert.Equal(t, "this", foundItem)
	assert.True(t, exists)

	t.Logf("found item: %s", foundItem)
}

func TestFindLast(t *testing.T) {
	var list = []string{"this", "test", "item"}
	var predicate = func(item string) bool {
		var chars = []rune(item)
		return chars[0] == 't'
	}

	var foundItem, exists = core.FindLast(list, predicate)

	assert.Equal(t, "test", foundItem)
	assert.True(t, exists)

	t.Logf("found item: %s", foundItem)
}

func TestKeys(t *testing.T) {
	var m = map[string]int{
		"this": 1,
		"test": 2,
		"item": 3,
	}

	var keys = core.Keys(m)
	sort.Strings(keys)

	assert.Equal(t, 3, len(keys))
	assert.Equal(t, keys[0], "item")
	assert.Equal(t, keys[1], "test")
	assert.Equal(t, keys[2], "this")

	t.Logf("keys: %v", keys)
}

func TestValues(t *testing.T) {
	var m = map[string]int{
		"this": 3,
		"test": 2,
		"item": 1,
	}

	var values = core.Values(m)
	sort.Ints(values)

	assert.Equal(t, 3, len(values))
	assert.Equal(t, values[0], 1)
	assert.Equal(t, values[1], 2)
	assert.Equal(t, values[2], 3)

	t.Logf("values: %v", values)
}

func TestEntries(t *testing.T) {
	var m = map[string]int{
		"this": 1,
		"test": 2,
		"item": 3,
	}

	var entries = core.Entries(m)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Key < entries[j].Key
	})

	assert.Equal(t, 3, len(entries))
	assert.Equal(t, entries[0].Key, "item")
	assert.Equal(t, entries[0].Value, 3)
	assert.Equal(t, entries[1].Key, "test")
	assert.Equal(t, entries[1].Value, 2)
	assert.Equal(t, entries[2].Key, "this")
	assert.Equal(t, entries[2].Value, 1)

	t.Logf("entries: %v", entries)
}

func TestToMap(t *testing.T) {
	var list = []string{"this", "test", "item"}
	var getKey = func(item string) string {
		var chars = []rune(item)
		return string(chars[:2])
	}

	var result = core.ToMap(list, getKey)

	assert.Equal(t, 3, len(result))
	assert.Equal(t, result["th"], "this")
	assert.Equal(t, result["te"], "test")
	assert.Equal(t, result["it"], "item")

	t.Logf("map: %v", result)
}

func TestFlatten(t *testing.T) {
	var nestedList = [][]string{
		{"this", "test"},
		{"item", "example"},
	}

	var flattened = core.Flatten(nestedList)

	assert.Equal(t, 4, len(flattened))
	assert.Equal(t, flattened[0], "this")
	assert.Equal(t, flattened[1], "test")
	assert.Equal(t, flattened[2], "item")
	assert.Equal(t, flattened[3], "example")

	t.Logf("flattened list: %v", flattened)
}

type ItemRequest struct {
	ID       string
	Quantity uint
}

type Item struct {
	ID       string
	OrderID  string
	Name     string
	Quantity uint
}

type Order struct {
	ID       string
	Customer string
	Items    []Item
}

func TestRelated(t *testing.T) {
	var orderList = []Order{
		{ID: "1", Customer: "Alice", Items: []Item{}},
		{ID: "2", Customer: "Bob", Items: []Item{}},
	}
	var itemList = []Item{
		{ID: "a1", OrderID: "1", Name: "Apple", Quantity: 2},
		{ID: "b1", OrderID: "1", Name: "Banana", Quantity: 3},
		{ID: "c1", OrderID: "2", Name: "Carrot", Quantity: 5},
	}
	var relate = func(order Order, item Item) (Order, bool) {
		if order.ID == item.OrderID {
			order.Items = append(order.Items, item)
			return order, true
		} else {
			return order, false
		}
	}

	var related = core.Relate(orderList, itemList, relate)

	assert.Equal(t, 2, len(related))
	assert.Equal(t, related[0].Customer, "Alice")
	assert.Equal(t, 2, len(related[0].Items))
	assert.Equal(t, "Apple", related[0].Items[0].Name)
	assert.Equal(t, "Banana", related[0].Items[1].Name)
	assert.Equal(t, related[1].Customer, "Bob")
	assert.Equal(t, 1, len(related[1].Items))
	assert.Equal(t, "Carrot", related[1].Items[0].Name)

	t.Logf("related items: %v", related)
}

func TestIntersect(t *testing.T) {
	var verticalList = []ItemRequest{
		{ID: "a1", Quantity: 2},
		{ID: "b1", Quantity: 3},
		{ID: "d1", Quantity: 4},
	}
	var horizontalList = []Item{
		{ID: "a1", OrderID: "1", Name: "Apple", Quantity: 1},
		{ID: "b1", OrderID: "1", Name: "Banana", Quantity: 1},
		{ID: "c1", OrderID: "2", Name: "Carrot", Quantity: 1},
	}
	var predicate = func(itemRequest ItemRequest, item Item) bool {
		return itemRequest.ID == item.ID
	}

	var verticalMatched, horizontalMatched, verticalUnMatched, horizontalUnMatched = core.Intersect(verticalList, horizontalList, predicate)

	assert.Equal(t, 2, len(horizontalMatched))
	assert.Equal(t, "Apple", horizontalMatched[0].Name)
	assert.Equal(t, "Banana", horizontalMatched[1].Name)
	assert.Equal(t, 2, len(verticalMatched))
	assert.Equal(t, uint(2), verticalMatched[0].Quantity)
	assert.Equal(t, uint(3), verticalMatched[1].Quantity)
	assert.Equal(t, 1, len(horizontalUnMatched))
	assert.Equal(t, "Carrot", horizontalUnMatched[0].Name)
	assert.Equal(t, 1, len(verticalUnMatched))
	assert.Equal(t, uint(4), verticalUnMatched[0].Quantity)

	t.Logf("verticalMatched: %v", verticalMatched)
	t.Logf("horizontalMatched: %v", horizontalMatched)
	t.Logf("verticalUnMatched: %v", verticalUnMatched)
	t.Logf("horizontalUnMatched: %v", horizontalUnMatched)
}

func TestGroup(t *testing.T) {
	var list = []Item{
		{ID: "a1", OrderID: "1", Name: "Apple", Quantity: 2},
		{ID: "b1", OrderID: "1", Name: "Banana", Quantity: 3},
		{ID: "c1", OrderID: "2", Name: "Carrot", Quantity: 2},
	}

	var grouper = func(item1 Item, item2 Item) bool {
		return item1.Quantity == item2.Quantity
	}

	var grouped = core.Group(list, grouper)

	assert.Equal(t, 2, len(grouped))
	assert.Equal(t, 2, len(grouped[0]))
	assert.Equal(t, 1, len(grouped[1]))
	assert.Equal(t, grouped[0][0].Name, "Apple")
	assert.Equal(t, grouped[0][1].Name, "Carrot")
	assert.Equal(t, grouped[1][0].Name, "Banana")

	t.Logf("grouped items: %v", grouped)
}

func TestDuplicates(t *testing.T) {
	var list = []string{"this", "test", "item", "test", "this"}
	var predicate = func(item1 string, item2 string) bool {
		return item1 == item2
	}

	var duplicates = core.Duplicate(list, predicate)

	assert.Equal(t, 4, len(duplicates))
	assert.Equal(t, duplicates[0], "this")
	assert.Equal(t, duplicates[1], "this")
	assert.Equal(t, duplicates[2], "test")
	assert.Equal(t, duplicates[3], "test")

	t.Logf("duplicates: %v", duplicates)
}
