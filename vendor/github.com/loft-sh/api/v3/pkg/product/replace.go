package product

import (
	"fmt"
	"strings"
)

// Replace replaces the product name in the given usage string
// based on the current product.Product().
//
// It replaces "loft" with the specific product name:
//   - "devpod pro" for product.DevPodPro
//   - "vcluster pro" for product.VClusterPro
//   - No replacement for product.Loft
//
// This handles case insensitive replaces like "loft" -> "devpod pro".
//
// It also handles case sensitive replaces:
//   - "Loft" -> "DevPod.Pro" for product.DevPodPro
//   - "Loft" -> "vCluster.Pro" for product.VClusterPro
//
// This allows customizing command usage text for different products.
//
// Parameters:
//   - content: The string to update
//
// Returns:
//   - The updated string with product name replaced if needed.
func Replace(content string) string {
	switch Product() {
	case DevPodPro:
		content = strings.Replace(content, "loft.com", "devpod.pro", -1)
		content = strings.Replace(content, "loft.host", "lfthost", -1)

		content = strings.Replace(content, "loft", "devpod pro", -1)
		content = strings.Replace(content, "Loft", "DevPod.Pro", -1)

		content = strings.Replace(content, "lfthost", "loft.host", -1)
	case VClusterPro:
		content = strings.Replace(content, "loft.com", "vcluster.pro", -1)
		content = strings.Replace(content, "loft.host", "lfthost", -1)

		content = strings.Replace(content, "loft", "vcluster pro", -1)
		content = strings.Replace(content, "Loft", "vCluster.Pro", -1)

		content = strings.Replace(content, "lfthost", "loft.host", -1)
	case Loft:
	}

	return content
}

// ReplaceWithHeader replaces the "loft" product name in the given
// usage string with the specific product name based on product.Product().
// It also adds a header with padding around the product name and usage.
//
// The product name replacements are:
//
//   - "devpod pro" for product.DevPodPro
//   - "vcluster pro" for product.VClusterPro
//   - No replacement for product.Loft
//
// This handles case insensitive replaces like "loft" -> "devpod pro".
//
// It also handles case sensitive replaces:
//   - "Loft" -> "DevPod.Pro" for product.DevPodPro
//   - "Loft" -> "vCluster.Pro" for product.VClusterPro
//
// Parameters:
//   - use: The usage string
//   - content: The content string to run product name replacement on
//
// Returns:
//   - The content string with product name replaced and header added
func ReplaceWithHeader(use, content string) string {
	maxChar := 56

	productName := "loft"

	switch Product() {
	case DevPodPro:
		productName = "devpod pro"
	case VClusterPro:
		productName = "vcluster pro"
	case Loft:
	}

	paddingSize := (maxChar - 2 - len(productName) - len(use)) / 2

	separator := strings.Repeat("#", paddingSize*2+len(productName)+len(use)+2+1)
	padding := strings.Repeat("#", paddingSize)

	return fmt.Sprintf(`%s
%s %s %s %s
%s
%s
`, separator, padding, productName, use, padding, separator, Replace(content))
}
