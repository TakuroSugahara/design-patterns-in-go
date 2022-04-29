package composite

import "fmt"

// Client
func Execute() {
	fmt.Printf("Execute composite pattern.\n")

	// --bin
	//   └--vi
	//   └--latex
	// --tmp
	// --usr
	//   ├--yuki
	//   | └--diary.html
	//   | └--foo.go
	//   └--hanako
	//     └--memo.txt

	vi := &file{name: "vi"}
	latex := &file{name: "latex"}
	bin := &folder{name: "bin"}
	bin.add(vi)
	bin.add(latex)

	bin.search("bin")
	fmt.Printf("###############\n")

	tmp := &folder{name: "tmp"}

	tmp.search("tmp")
	fmt.Printf("###############\n")

	usr := &folder{name: "usr"}
	yuki := &folder{name: "yuki"}
	hanako := &folder{name: "hanako"}

	usr.add(yuki)
	usr.add(hanako)

	diary := &file{name: "diary.html"}
	foo := &file{name: "foo.go"}
	yuki.add(diary)
	yuki.add(foo)

	memo := &file{name: "memo.txt"}
	hanako.add(memo)

	usr.search("usr")
	fmt.Printf("###############\n")

	fmt.Printf("Finished composite pattern.\n\n")
}
