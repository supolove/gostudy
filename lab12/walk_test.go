package lab12

import (
	"gostudy/lab12/filepath"
	"gostudy/lab12/pet"
	"testing"
)

func TestWalk(t *testing.T) {
	var petWalker pet.Walker
	var filepathWalker filepath.Walker
	println("call by pet.Walker:")
	petWalker = &pet.Dog{}
	petWalker.Walk()
	petWalker = &pet.Cat{}
	petWalker.Walk()
	petWalker = &filepath.FilePath{}
	petWalker.Walk()

	println("\ncall by filepath.Walker:")
	filepathWalker = &pet.Dog{}
	filepathWalker.Walk()
	filepathWalker = &pet.Cat{}
	filepathWalker.Walk()
	filepathWalker = &filepath.FilePath{}
	filepathWalker.Walk()

	println("\nThe strange is that pet.Warker and filepath.Walker has the same signiture but they are not the same one.")
	println("But Go treat them as the same one.")
	println(`
implements pet.Walker{
	*pet.Dog
	*pet.Cat
}
implements filepath.Walker{
	*filepath.FilePath
}
`)
}
