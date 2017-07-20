package graphql_examples

import (
	"encoding/json"
	"net/url"
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
	b, err := get("http://localhost:8080/graphql?query=query{hello(message:\"World\")}")
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
	b, err := get("http://localhost:8080/graphql?query=query{bye}")
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
	b, err := get("http://localhost:8080/graphql?query=mutation{sum(x:1,y:2)}")
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

func urlencode(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func TestHttp_User(t *testing.T) {
	result := new(struct {
		Data struct {
			User struct {
				Name string
			}
		}
	})

	b, err := get("http://localhost:8080/graphql?query=query{user(id:\"1\"){name}}")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, json.Unmarshal(b, result))
	assert.Equal(t, result.Data.User.Name, "Name-1")

	b, err = get("http://localhost:8080/graphql?query=query{user(id:\"id\"){name}}")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, json.Unmarshal(b, result))
	assert.Equal(t, result.Data.User.Name, "Name")

	chinese, err := urlencode("\"编号\"")
	if err != nil {
		t.Fatal(err)
	}
	b, err = get("http://localhost:8080/graphql?query=query{user(id:" + chinese + "){name}}")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, json.Unmarshal(b, result))
	assert.Equal(t, result.Data.User.Name, "名字")
}
