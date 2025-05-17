package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shorter/internal/logic"
	"shorter/internal/svc"
)

func healthzHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHealthzLogic(r.Context(), svcCtx)
		err := l.Healthz()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
