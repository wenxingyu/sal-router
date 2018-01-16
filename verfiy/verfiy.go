package verfiy

import "github.com/wenxingyu/sal-router/model"

type Validate interface {
	do(request model.UserRequest) bool
}
