package main

type Invocation struct {
	Name           string         `@Ident`
	Arguments      *[]Value       `"(" ( @Value )* ")"`
	NamedArguments *NamedArgument `| "(" @@* ")"`
}

type NamedArgument struct {
	Name  string `@Ident :`
	Value Value  `@@`
}

type Value struct {
	String *string  `  @String`
	Float  *float64 `| @Float`
	Int    *int     `| @Int`
}
