package graphql

import (
	"odunlamizo/book-collection/internal/database"
	"odunlamizo/book-collection/internal/model"

	"github.com/graphql-go/graphql"
)

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"co_authors": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"SearchByTitle": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get book(s) by title",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					title, ok := p.Args["title"].(string)
					var (
						books []model.Book
						err   error
					)
					if ok {
						books, err = database.GetBooksByTitle(title)
						if err != nil {
							return nil, err
						}
					}
					return books, nil
				},
			},
			"SearchByAuthor": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get book(s) by author",
				Args: graphql.FieldConfigArgument{
					"author": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					author, ok := p.Args["author"].(string)
					var (
						books []model.Book
						err   error
					)
					if ok {
						books, err = database.GetBooksByAuthor(author)
						if err != nil {
							return nil, err
						}
					}
					return books, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"Add": &graphql.Field{
				Type:        bookType,
				Description: "Add new book",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"url": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"co_authors": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					book := &model.Book{
						Title:  params.Args["title"].(string),
						Url:    params.Args["url"].(string),
						Author: params.Args["author"].(string),
						CoAuthors: func() []string {
							coAuthorsRaw, ok := params.Args["co_authors"].([]interface{})
							if ok {
								coAuthors := make([]string, len(coAuthorsRaw))
								for index, author := range coAuthorsRaw {
									coAuthors[index] = author.(string)
								}
								return coAuthors
							} else {
								return []string{}
							}
						}(),
					}
					if err := database.AddBook(book); err != nil {
						return nil, err
					}
					return *book, nil
				},
			},
			"DeleteById": &graphql.Field{
				Type:        bookType,
				Description: "Delete book by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(int)
					var (
						book *model.Book
						err  error
					)
					if ok {
						book, err = database.DeleteBookById(id)
						if err != nil {
							return nil, err
						}
					}
					return book, nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
