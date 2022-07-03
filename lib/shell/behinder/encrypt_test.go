package behinder

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_decryptForCSharp(t *testing.T) {
	// rNA7qt3vG0zX1iz0oKDWMxkXILuL519QB3HK8qDzThNHCJVpsqHACEF1Hii2t7NV2EeWa8rSqygShPuOzAVrbdDlEfe1aQzNCm
	bs, _ := base64.StdEncoding.DecodeString("vG7dOhhPqMoYV6WC3Rj3Fte/Mnqq2agRfeSGMrdO5Na+3IC4Id1mguhqCs16rNt1yP7RXh4/5IPY4oW29UJ8w4ZsQwr+COJcso0A8GOeUCv9tVNgDSWudEb1eqtXgRgfnmSTrnSRkwJCocHy8a4q6qqQ5dmtr+77GsjOjQe+BIlF2KFh6c8gJQAzpQGduRhmy9awATimy4lIENSWl/zWn6pKh2ZdWMPPh/JQKRcPBIA=")
	key := []byte("e45e329feb5d925b")
	a := decryptForCSharp(bs, key)
	fmt.Printf("%#+v\n", string(a))
	bs1, _ := base64.StdEncoding.DecodeString("pz3NtQ2T9gBozbTYUyAeE29hGq+F+f44WaO+xFDIXOKsYVGp4Qa6q4wPHrocKxQRxrXVImAurvvMXKTiRhC9TmKG5yCDtC8Zv6+eU4rt+k1tAooWUIPMT/HQTQFrrljdHjZDHRYxT4i3mNBw2pl3Lncd8nFPDItU3ScAuCAAAYLNvLRauz6I8PLcX3QKNaZ2EXEpmtMM+2f5YSugGhbfgyJgGwjiZgqB2CKafvyDcJHozUhcEB7c4sGfoN+SgxqF")
	key1 := []byte("e45e329feb5d925b")
	a1 := decryptForCSharp(bs1, key1)
	fmt.Printf("%#+v\n", string(a1))
	bs2, _ := base64.StdEncoding.DecodeString("vG7dOhhPqMoYV6WC3Rj3Fte/Mnqq2agRfeSGMrdO5Na+3IC4Id1mguhqCs16rNt1yP7RXh4/5IPY4oW29UJ8w4ZsQwr+COJcso0A8GOeUCv9tVNgDSWudEb1eqtXgRgfnmSTrnSRkwJCocHy8a4q6qqQ5dmtr+77GsjOjQe+BIlF2KFh6c8gJQAzpQGduRhmy9awATimy4lIENSWl/zWn20iKHATK0bzGpm3lDU4fHk=")
	key2 := []byte("e45e329feb5d925b")
	a2 := decryptForCSharp(bs2, key2)
	fmt.Printf("%#+v\n", string(a2))
}
