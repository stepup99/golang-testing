package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/stepup99/golang-testing/user_gin/domain"
)

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetUser(userId int) (*domain.User, error) {
	fmt.Println("userid >>>>>>> inside service ", userId)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%v", userId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("userid >>>>>>> inside service 29 ")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user %s", resp.Status)
	}
	fmt.Println("userid >>>>>>> inside service 39 ")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user domain.User
	fmt.Println("userid >>>>>>> inside service 46 ")
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	fmt.Println("userid >>>>>>> inside service 51 ")
	return &user, nil
}
