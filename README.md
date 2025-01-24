# Go Nil Map Panic

This example demonstrates a common error in Go: accessing a key in a nil map.  Go maps are reference types, and if not initialized they are `nil`. Attempting to access a key in a `nil` map will result in a runtime panic. This is a subtle error that is often missed during testing.