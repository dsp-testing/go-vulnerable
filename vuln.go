package go_vulnerable

import "github.com/ulikunitz/xz"

func Do() {
	_ = xz.Reader{}
}
