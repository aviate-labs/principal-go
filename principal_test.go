package principal_test

import (
	"encoding/hex"
	"fmt"

	principal "github.com/allusion-be/principal-go"
)

func ExampleDecode() {
	p, _ := principal.Decode("em77e-bvlzu-aq")
	fmt.Printf("%x", p)
	// Output:
	// abcd01
}

func ExamplePrincipal() {
	raw, _ := hex.DecodeString("abcd01")
	p := principal.Principal(raw)
	fmt.Println(p.Encode())
	// Output:
	// em77e-bvlzu-aq
}
