package middleware

const StopServingStatus = 0

const (
	NotNeedUpdate = "0" //不需要更新
	WeakUpdate    = "1" //弱更新
	StrongUpdate  = "2" //强更新
)

//func TokenAuth() echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(ctx echo.Context) (err error) {
//			req := new(pblogin.VerifyUserTokenReq)
//			req.Token, err = url.PathUnescape(ctx.Request().Header.Get("Authorization"))
//			context := opentracing.ContextFromEcho(ctx)
//			resp, err := rpc_client.LoginClient.VerifyUserToken(context, req)
//			if err != nil {
//				return httputil.ErrorResponseJson(ctx, err)
//			}
//			ctx.Set("user_id", resp.UserId)
//			ctx.Response().Header().Set("Authorization", "Bearer "+resp.Token)
//
//			return next(ctx)
//		}
//	}
//}
//
////极验行为验证获取登录态
//func GetLoginStatus() echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(ctx echo.Context) (err error) {
//			req := new(pblogin.VerifyUserTokenReq)
//			req.Token, err = url.PathUnescape(ctx.Request().Header.Get("Authorization"))
//			context := opentracing.ContextFromEcho(ctx)
//			resp, err := rpcclient.LoginClient.VerifyUserToken(context, req)
//			var userId int64
//			if err != nil {
//				userId = 0
//			} else {
//				userId = resp.UserId
//				ctx.Response().Header().Set("Authorization", "Bearer "+resp.Token)
//			}
//			ctx.Set("user_id", userId)
//			return next(ctx)
//		}
//	}
//}
//
//func StatusAuth() echo.MiddlewareFunc {
//	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
//		return func(context echo.Context) error {
//			req := new(pbsystemcache.Req)
//			ctx := opentracing.ContextFromEcho(context)
//
//			resp, err := rpcclient.SystemCacheClient.GetSystemStatusInfo(ctx, req)
//			if err != nil {
//				return httputil.ErrorResponseJson(context, err)
//			}
//
//			for _, value := range resp.SystemStatus {
//				if value.Status == StopServingStatus {
//					return httputil.ErrorResponseJson(context, unicornstd.ErrSystemStopServing)
//				}
//			}
//
//			return handlerFunc(context)
//
//		}
//	}
//}
