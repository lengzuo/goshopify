package common

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	shopifyDomain    = "myshopify.com"
	dotShopifyDomain = "." + shopifyDomain
	shopScheme       = "https"
)

// ShopDomain input shop name and get the shopDomain like <shop>.myshopify.com
func ShopDomain(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Trim(name, ".")
	if strings.Contains(name, shopifyDomain) {
		return name
	}
	return name + dotShopifyDomain
}

// ShopShortName return the short shop name without .myshopify.com
func ShopShortName(name string) string {
	return strings.Replace(ShopDomain(name), dotShopifyDomain, "", -1)
}

// ShopBaseUrl return the Shop's base url.
func ShopBaseUrl(name string) string {
	return fmt.Sprintf("%s://%s", shopScheme, ShopDomain(name))
}

// Int64Str use to convert int64 to string
// it is more efficient than we do fmt.Sprintf("%d",v)
func Int64Str(v int64) string {
	return strconv.FormatInt(v, 10)
}
