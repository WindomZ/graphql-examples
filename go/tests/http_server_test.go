package graphql_examples

import (
	"encoding/json"
	"testing"

	"github.com/WindomZ/go-develop-kit/http"
	"github.com/WindomZ/testify/assert"
)

var client *http.HttpClient = http.NewHttpClient()

func get(uri string) ([]byte, error) {
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	return http.ByteResponseBody(resp)
}

func TestHttp_Hello(t *testing.T) {
	b, err := get("http://localhost:8080/example?graphql=query{hello(message:\"World\")}")
	if err != nil {
		t.Fatal(err)
	}

	result := new(struct {
		Data struct {
			Hello string
		}
	})
	assert.NoError(t, json.Unmarshal(b, result))

	assert.Equal(t, result.Data.Hello, "Hello World!")
}

func TestHttp_Bye(t *testing.T) {
	b, err := get("http://localhost:8080/example?graphql=query{bye}")
	if err != nil {
		t.Fatal(err)
	}

	result := new(struct {
		Data struct {
			Bye string
		}
	})
	assert.NoError(t, json.Unmarshal(b, result))

	assert.Equal(t, result.Data.Bye, "Goodbye!")
}

func TestHttp_Sum(t *testing.T) {
	b, err := get("http://localhost:8080/example?graphql=mutation{sum(x:1,y:2)}")
	if err != nil {
		t.Fatal(err)
	}

	result := new(struct {
		Data struct {
			Sum int
		}
	})
	assert.NoError(t, json.Unmarshal(b, result))

	assert.Equal(t, result.Data.Sum, 3)
}

func TestHttp_User(t *testing.T) {
	result := new(struct {
		Data struct {
			User struct {
				Name string
			}
		}
	})

	b, err := get("http://localhost:8080/example?graphql=query{user(id:\"1\"){name}}")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, json.Unmarshal(b, result))
	assert.Equal(t, result.Data.User.Name, "Name-1")

	b, err = get("http://localhost:8080/example?graphql=query{user(id:\"id\"){name}}")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, json.Unmarshal(b, result))
	assert.Equal(t, result.Data.User.Name, "Name")

	b, err = get("http://localhost:8080/example?graphql=query{user(id:\"编号\"){name}}")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, json.Unmarshal(b, result))
	assert.Equal(t, result.Data.User.Name, "名字")
}
