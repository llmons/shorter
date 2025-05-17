package handler

import (
	"net/http"

	"shorter/internal/logic"
	"shorter/internal/svc"
	"shorter/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func redirectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqShortCode
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRedirectLogic(r.Context(), svcCtx)
		resp, err := l.Redirect(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			w.Header().Set("Location", resp.Url)
			w.WriteHeader(http.StatusMovedPermanently)
		}
	}
}
