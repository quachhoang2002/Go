package http

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quachhoang2002/Go/pkg/response"
)

func (h handler) add(c *gin.Context) {
	ctx := c.Request.Context()
	var req addRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errInvalidRequestBody)
		return
	}

	fmt.Println(req)

	input, err := req.toInput()
	if err != nil {
		response.Error(c, err)
		return
	}

	if err := h.accountUC.Create(ctx, input); err != nil {
		h.l.Errorf(ctx, "account.handler.add.accountUC.Create(ctx, %+v): %s", input, err)
		response.ErrorWithMap(c, err, errMap)
		return
	}

	response.OK(c, nil)
}

func (h handler) list(c *gin.Context) {
	ctx := c.Request.Context()

	todos, err := h.accountUC.All(ctx)
	if err != nil {
		h.l.Error(ctx, "todo.handler.list.todoUC.All: %s", err)
		response.ErrorWithMap(c, err, errMap)
		return
	}

	response.OK(c, newListResponse(todos))
}

func (h handler) delete(c *gin.Context) {
	ctx := c.Request.Context()

	rawID := c.Param("id")
	if strings.TrimSpace(rawID) == "" {
		response.Error(c, errIDIsRequired)
		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		h.l.Warnf(ctx, "todo.handler.delete.strconv.Atoi(%s): %s", rawID, err)
		response.Error(c, errInvalidID)
		return
	}

	if err := h.accountUC.Delete(ctx, id); err != nil {
		h.l.Errorf(ctx, "todo.handler.delete.todoUC.Delete(ctx, %d): %s", id, err)
		response.ErrorWithMap(c, err, errMap)
		return
	}

	response.OK(c, nil)
}
