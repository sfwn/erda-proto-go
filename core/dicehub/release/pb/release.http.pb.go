// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: release.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	http1 "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// ReleaseServiceHandler is the server API for ReleaseService service.
type ReleaseServiceHandler interface {
	// POST /api/releases
	CreateRelease(context.Context, *ReleaseCreateRequest) (*ReleaseCreateResponseData, error)
	// PUT /api/releases/{releaseID}
	UpdateRelease(context.Context, *ReleaseUpdateRequest) (*ReleaseDataResponse, error)
	// PUT /api/releases/{releaseID}/reference/actions/change
	UpdateReleaseReference(context.Context, *ReleaseReferenceUpdateRequest) (*ReleaseDataResponse, error)
	// GET /api/releases/{releaseID}/actions/get-plist
	GetIosPlist(context.Context, *GetIosPlistRequest) (*GetIosPlistResponse, error)
	// GET /api/releases/{releaseID}
	GetRelease(context.Context, *GetIosPlistRequest) (*ReleaseGetResponse, error)
	// DELETE /api/releases/{releaseID}
	DeleteRelease(context.Context, *GetIosPlistRequest) (*ReleaseDataResponse, error)
	// GET /api/releases
	ListRelease(context.Context, *ReleaseListRequest) (*ReleaseListResponse, error)
	// GET /api/releases/actions/get-name
	ListReleaseName(context.Context, *ListReleaseNameRequest) (*ListReleaseNameResponse, error)
	// GET /api/releases/actions/get-latest
	GetLatestReleases(context.Context, *GetLatestReleasesRequest) (*GetLatestReleasesResponse, error)
	// POST /gc
	ReleaseGC(context.Context, *ReleaseGCRequest) (*ReleaseDataResponse, error)
}

// RegisterReleaseServiceHandler register ReleaseServiceHandler to http.Router.
func RegisterReleaseServiceHandler(r http.Router, srv ReleaseServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_CreateRelease := func(method, path string, fn func(context.Context, *ReleaseCreateRequest) (*ReleaseCreateResponseData, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseCreateRequest))
		}
		var CreateRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "CreateRelease", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in ReleaseCreateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateRelease_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateRelease := func(method, path string, fn func(context.Context, *ReleaseUpdateRequest) (*ReleaseDataResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseUpdateRequest))
		}
		var UpdateRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "UpdateRelease", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in ReleaseUpdateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateRelease_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateReleaseReference := func(method, path string, fn func(context.Context, *ReleaseReferenceUpdateRequest) (*ReleaseDataResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseReferenceUpdateRequest))
		}
		var UpdateReleaseReference_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateReleaseReference_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "UpdateReleaseReference", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in ReleaseReferenceUpdateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateReleaseReference_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetIosPlist := func(method, path string, fn func(context.Context, *GetIosPlistRequest) (*GetIosPlistResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetIosPlistRequest))
		}
		var GetIosPlist_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetIosPlist_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "GetIosPlist", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GetIosPlistRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetIosPlist_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetRelease := func(method, path string, fn func(context.Context, *GetIosPlistRequest) (*ReleaseGetResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetIosPlistRequest))
		}
		var GetRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "GetRelease", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GetIosPlistRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetRelease_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DeleteRelease := func(method, path string, fn func(context.Context, *GetIosPlistRequest) (*ReleaseDataResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetIosPlistRequest))
		}
		var DeleteRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "DeleteRelease", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GetIosPlistRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteRelease_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListRelease := func(method, path string, fn func(context.Context, *ReleaseListRequest) (*ReleaseListResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseListRequest))
		}
		var ListRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ListRelease", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in ReleaseListRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListRelease_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListReleaseName := func(method, path string, fn func(context.Context, *ListReleaseNameRequest) (*ListReleaseNameResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListReleaseNameRequest))
		}
		var ListReleaseName_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListReleaseName_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ListReleaseName", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in ListReleaseNameRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListReleaseName_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetLatestReleases := func(method, path string, fn func(context.Context, *GetLatestReleasesRequest) (*GetLatestReleasesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetLatestReleasesRequest))
		}
		var GetLatestReleases_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetLatestReleases_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "GetLatestReleases", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GetLatestReleasesRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetLatestReleases_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ReleaseGC := func(method, path string, fn func(context.Context, *ReleaseGCRequest) (*ReleaseDataResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseGCRequest))
		}
		var ReleaseGC_info transport.ServiceInfo
		if h.Interceptor != nil {
			ReleaseGC_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ReleaseGC", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in ReleaseGCRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ReleaseGC_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateRelease("POST", "/api/releases", srv.CreateRelease)
	add_UpdateRelease("PUT", "/api/releases/{releaseID}", srv.UpdateRelease)
	add_UpdateReleaseReference("PUT", "/api/releases/{releaseID}/reference/actions/change", srv.UpdateReleaseReference)
	add_GetIosPlist("GET", "/api/releases/{releaseID}/actions/get-plist", srv.GetIosPlist)
	add_GetRelease("GET", "/api/releases/{releaseID}", srv.GetRelease)
	add_DeleteRelease("DELETE", "/api/releases/{releaseID}", srv.DeleteRelease)
	add_ListRelease("GET", "/api/releases", srv.ListRelease)
	add_ListReleaseName("GET", "/api/releases/actions/get-name", srv.ListReleaseName)
	add_GetLatestReleases("GET", "/api/releases/actions/get-latest", srv.GetLatestReleases)
	add_ReleaseGC("POST", "/gc", srv.ReleaseGC)
}
