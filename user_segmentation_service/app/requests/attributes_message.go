package requests

var Attributes = make(map[string]string)

func initAttributes() {
	Attributes["Name"] = "name"
	Attributes["Config"] = "config"
	Attributes["Images"] = "images"
	Attributes["KitBlockId"] = "Kit Block"
	Attributes["Pieces"] = "pieces"
	Attributes["Visibility"] = "visibility"
	Attributes["Property"] = "property"
	Attributes["Material"] = "material"
	Attributes["Page"] = "page"
	Attributes["Limit"] = "limit"
	Attributes["Mine"] = "mine"
	Attributes["OrderColumn"] = "order column"
	Attributes["OrderType"] = "order type"
	Attributes["Id"] = "id"
	Attributes["Owner"] = "owner"
	Attributes["Wallet"] = "wallet"
	Attributes["File"] = "file"
	Attributes["Description"] = "description"
	Attributes["Style"] = "style"
	Attributes["Query"] = "query"
}
