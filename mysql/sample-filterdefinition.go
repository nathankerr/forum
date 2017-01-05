package mysql

func main() {

	// Simple Equals Filter for users Table
	equalsFilter := FilterDefinition{FilterType: Equals, FieldName: "username", Values: []string{"jeremy", "dominique"}}

	// And with 2 Equals Filter for users Table
	andEqualsFilter := FilterDefinition{FilterType: And, Items: []FilterDefinition{
		FilterDefinition{FilterType: Equals, FieldName: "username", Values: []string{"jeremy", "dominique"}},
		FilterDefinition{FilterType: Equals, FieldName: "password", Values: []string{"test123"}}}}
}
