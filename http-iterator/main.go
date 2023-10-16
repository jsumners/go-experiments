package httpiter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type page struct {
	Page       int
	PerPage    int `json:"per_page"`
	Total      int
	TotalPages int `json:"total_pages"`
	Data       []User
}

type User struct {
	Id        int
	Email     string
	FirstName string
	LastName  string
	Avatar    string
}

type UsersIter struct {
	current *page
}

var Done = errors.New("done")

func NewUsersIter() *UsersIter {
	result := &UsersIter{
		current: &page{},
	}

	return result
}

func getPage(pageNum int) (*http.Response, error) {
	getUrl := url.URL{
		Scheme: "https",
		Host:   "reqres.in",
		Path:   "/api/users",
	}

	query := getUrl.Query()
	query.Add("per_page", "3")
	query.Add("page", fmt.Sprintf("%d", pageNum))
	getUrl.RawQuery = query.Encode()

	return http.Get(getUrl.String())
}

func (ui *UsersIter) Next() (*[]User, error) {
	if (ui.current.Page > 0) && (ui.current.Page == ui.current.TotalPages) {
		return nil, Done
	}

	res, err := getPage(ui.current.Page + 1)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var _page page
	err = json.NewDecoder(res.Body).Decode(&_page)
	if err != nil {
		return nil, err
	}

	*ui.current = _page
	return &ui.current.Data, nil
}
