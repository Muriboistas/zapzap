package mapx

import "fmt"

// FromTo change a key to another map
func FromTo(key string, from, to *map[string]string) error {
	v, found := (*from)[key]
	if !found {
		return fmt.Errorf("key %q not founded", key)
	}
	(*to)[key] = v
	delete(*from, key)

	return nil
}
