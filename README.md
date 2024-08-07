# Go source code obfuscator

The goal with this project is to provide a go tool that can take a go
repository and obfuscate all "user defined" symbols. The resulting, obfuscated
repository has to be compilable. *Correctly* obfuscating tests is not
supported.

So something like

```go
func main() {
	start := time.Now()
	moduleNameFlag := flag.String("moduleName", "", "The name of the module (top of go.mod).")
	sourceFlag := flag.String("source", "", "The full path of the source repository.")
	sinkFlag := flag.String("sink", "", "The full path where to write obfuscated directory.")
	flag.Parse()

	args, err := args.NewArgs(moduleNameFlag, sourceFlag, sinkFlag)
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("%v", err)
	}

	if err := run(args); err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Run took: %v", time.Since(start))
}
```

becomes

```go
func main() {
	ccedcicgdcdfjjccdjgafeaadhdcbfgbcibecddfijcbbcdaegafgcdbfdffccfaXXX := time.Now()
	dfedbdicafbbfbebdafajahgdccdbeadcdbgabedjdabibefaaecjgcdadcejcdeXXX := flag.String("moduleName", "", "The name of the module (top of go.mod).")
	ecjficcjedfceagfbcfbchfibbacjbbecbjhaabecgaejeggbhcebhebgijcifbfXXX := flag.String("source", "", "The full path of the source repository.")
	ajcdiacfcfegafdijaeagdjdagfiahhigacfbeafgcdebcgiicjdgcacbfjhciccXXX := flag.String("sink", "", "The full path where to write obfuscated directory.")
	flag.Parse()

	ajahhccfeagijhddaaddfhbfebhiiddjfeccacbeecefdigdeibfaefjbjidgcdhXXX, djebcfdeagjihfaheafddbijfhddjfhaibaeccahigfafbbfcbcjfffjhcffehhfXXX := ajahhccfeagijhddaaddfhbfebhiiddjfeccacbeecefdigdeibfaefjbjidgcdhXXX.NdfbfghgaeciabcjcdcifecdcgcjededeicabcfccfjhceabbafcafbaehgabaeaXXX(dfedbdicafbbfbebdafajahgdccdbeadcdbgabedjdabibefaaecjgcdadcejcdeXXX, ecjficcjedfceagfbcfbchfibbacjbbecbjhaabecgaejeggbhcebhebgijcifbfXXX, ajcdiacfcfegafdijaeagdjdagfiahhigacfbeafgcdebcgiicjdgcacbfjhciccXXX)
	if djebcfdeagjihfaheafddbijfhddjfhaibaeccahigfafbbfcbcjfffjhcffehhfXXX != nil {
		flag.PrintDefaults()
		log.Fatalf("%v", djebcfdeagjihfaheafddbijfhddjfhaibaeccahigfafbbfcbcjfffjhcffehhfXXX)
	}

	if djebcfdeagjihfaheafddbijfhddjfhaibaeccahigfafbbfcbcjfffjhcffehhfXXX := acbacffbcbaafiabfgfcdccdbecgfbeffdjeiaacdahhfifcffebaiihedjifjbeXXX(ajahhccfeagijhddaaddfhbfebhiiddjfeccacbeecefdigdeibfaefjbjidgcdhXXX); djebcfdeagjihfaheafddbijfhddjfhaibaeccahigfafbbfcbcjfffjhcffehhfXXX != nil {
		log.Fatalf("%v", djebcfdeagjihfaheafddbijfhddjfhaibaeccahigfafbbfcbcjfffjhcffehhfXXX)
	}
	log.Printf("Run took: %v", time.Since(ccedcicgdcdfjjccdjgafeaadhdcbfgbcibecddfijcbbcdaegafgcdbfdffccfaXXX))
}
```

As you can see, strings are not supported (yet(?)).

## Run

To run, provide the `sink` and `source` directories. The `source` is the
absolute path to the repository that should be ofuscated. Only repositories
with a go.mod file are supported. The `sink` is the
directory where the obfuscated repository will be written to.

e.g.

`go run main.go --source=/home/ada/repos/gosrcobfsc --sink=/home/ada/repos/gosrcobfsc/tests`

## Current State

Only this repository was tested and successfully obfuscated, so there are
guaranteed a lot of edge cases not covered.

### TODO

See [TODO](TODO.md)

