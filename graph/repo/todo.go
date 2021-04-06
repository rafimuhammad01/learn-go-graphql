package repo

import (
	"context"
	dbcon "github.com/rafimuhammad01/learn-go-graphql.git/graph/database"
	"github.com/rafimuhammad01/learn-go-graphql.git/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Todo struct {
	col *mongo.Collection
}

func NewTodo() *Todo {
	con := dbcon.Connect()

	col := con.Database.Collection("todo")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	return &Todo{col: col}
}

func (s *Todo) Save (input *model.NewTodo) *model.Todo {
	collection := s.col
	ctx := context.Background()

	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Todo{
		ID:       res.InsertedID.(primitive.ObjectID).Hex(),
		Activity: input.Activity,
		Detail:   input.Detail,
	}
}

func (s *Todo) Create (input *model.NewTodo) *model.Todo {
	ctx := context.Background()

	res, err := s.col.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Todo{
		ID:       res.InsertedID.(primitive.ObjectID).Hex(),
		Activity: input.Activity,
		Detail:   input.Detail,
	}
}

func (s *Todo) FindByID (ID string) *model.Todo {
	ctx := context.Background()
	objectID, err:= primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}

	res := s.col.FindOne(ctx, bson.M{"_id" : objectID})

	todo := model.Todo{}
	res.Decode(&todo)

	return &todo
}

func (s *Todo) FindAll () []*model.Todo{
	ctx := context.Background()

	res, err := s.col.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var todos []*model.Todo

	for res.Next(ctx) {
		var todo *model.Todo
		err := res.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	return todos
}

func (s *Todo) Update (id string, input *model.NewTodo) *model.Todo {
	ctx := context.Background()

	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	update := bson.D{
		{"$set", bson.D{{"activity", input.Activity}, {"detail", input.Detail}}, },
	}
	_, err = s.col.UpdateOne(ctx, bson.M{"_id" : ObjectID}, update)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Todo{
		ID:       id,
		Activity: input.Activity,
		Detail:   input.Detail,
	}
}

func (s *Todo) Delete (ID string) *model.Todo {
	ctx := context.Background()

	objectID, err:= primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	res := s.col.FindOne(ctx, bson.M{"_id" : objectID})

	_, err = s.col.DeleteOne(ctx, bson.M{"_id":objectID})
	if err != nil {
		log.Fatal(err)
	}

	todo := model.Todo{}
	res.Decode(&todo)

	return &todo
}

