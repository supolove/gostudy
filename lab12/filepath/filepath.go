package filepath

type FilePath struct {
}

func (fp *FilePath) Walk() {
	println("  filepath.FilePath.Walk")
}

type Walker interface {
	Walk()
}
