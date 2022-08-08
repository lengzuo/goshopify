package common

import "fmt"

type Collection string

var (
	Articles      = register("articles")
	Blogs         = register("blogs")
	Customers     = register("customers")
	Orders        = register("orders")
	Products      = register("products")
	Images        = register("images")
	Variants      = register("variants")
	ScriptTags    = register("script_tags")
	Shops         = register("shops")
	Theme         = register("themes")
	allCollection = map[string]Collection{}
)

func register(value string) Collection {
	if _, ok := allCollection[value]; ok {
		panic(fmt.Sprintf("collection value of %s registered", value))
	}
	allCollection[value] = Collection(value)
	return allCollection[value]
}

func (c Collection) String() string {
	return string(c)
}

func (c Collection) IsValid() bool {
	_, ok := allCollection[c.String()]
	return ok
}
