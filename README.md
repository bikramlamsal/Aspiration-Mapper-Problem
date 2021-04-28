# Aspiration-Mapper-Problem

mapper.go
/*
	1 . Write a function with this signature:
	
	func CapitalizeEveryThirdAlphanumericChar(s string) string {
		// your code goes here
	}
	
	that capitalizes *only* every third alphanumeric character of a string, so that if I hand you
	Aspiration.com
	You hand me back
	asPirAtiOn.cOm
	Please note: 
	- Characters other than each third should be downcased
	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9]. 
	2. Do the same problem, but this time create a "mapper" package that looks like this:
*/

package mapper

type Interface interface {
   TransformRune(pos int)
   GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
   for pos, _ := range i.GetValueAsRuneSlice() {
      i.TransformRune(pos)
   }
}

// And change your code to look like this:

func main() {
   s := NewSkipString(3, "Aspiration.com")
   mapper.MapString(&s)
   fmt.Println(s)
}

/*
  Make sure your fmt.Println(s) output looks nice and clean, ie implement the Stringer interface.
  Zip up your file(s) and email them to the recruiter
*/
