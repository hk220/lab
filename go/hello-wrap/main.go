// https://qiita.com/shibukawa/items/e633e426a6e67ea2e830
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("cause")
	err = fmt.Errorf("first: %w", err)
	err = fmt.Errorf("second: %w", err)
	err = fmt.Errorf("third: %w", err)
	// third: second: first: cause
	fmt.Println(err)
	// second: first: cause
	fmt.Println(errors.Unwrap(err))
}
