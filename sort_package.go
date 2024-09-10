package main

func sortPackage(width int, height int, length int, mass int) string {
	p := &Parcel{width, height, length, mass}
	return p.Stack()
}
