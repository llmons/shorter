package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shorter/internal/logic"
	"shorter/internal/svc"
	"shorter/internal/types"
)

func shortHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUrl
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShortLogic(r.Context(), svcCtx)
		resp, err := l.Short(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
