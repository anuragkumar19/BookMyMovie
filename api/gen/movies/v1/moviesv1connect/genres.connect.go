// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: movies/v1/genres.proto

package moviesv1connect

import (
	v1 "bookmymovie.app/bookmymovie/api/gen/movies/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MoviesGenresServiceName is the fully-qualified name of the MoviesGenresService service.
	MoviesGenresServiceName = "movies.v1.MoviesGenresService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MoviesGenresServiceGetGenreProcedure is the fully-qualified name of the MoviesGenresService's
	// GetGenre RPC.
	MoviesGenresServiceGetGenreProcedure = "/movies.v1.MoviesGenresService/GetGenre"
	// MoviesGenresServiceGetGenresProcedure is the fully-qualified name of the MoviesGenresService's
	// GetGenres RPC.
	MoviesGenresServiceGetGenresProcedure = "/movies.v1.MoviesGenresService/GetGenres"
	// MoviesGenresServiceCreateGenreProcedure is the fully-qualified name of the MoviesGenresService's
	// CreateGenre RPC.
	MoviesGenresServiceCreateGenreProcedure = "/movies.v1.MoviesGenresService/CreateGenre"
	// MoviesGenresServiceUpdateGenreProcedure is the fully-qualified name of the MoviesGenresService's
	// UpdateGenre RPC.
	MoviesGenresServiceUpdateGenreProcedure = "/movies.v1.MoviesGenresService/UpdateGenre"
	// MoviesGenresServiceDeleteGenreProcedure is the fully-qualified name of the MoviesGenresService's
	// DeleteGenre RPC.
	MoviesGenresServiceDeleteGenreProcedure = "/movies.v1.MoviesGenresService/DeleteGenre"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	moviesGenresServiceServiceDescriptor           = v1.File_movies_v1_genres_proto.Services().ByName("MoviesGenresService")
	moviesGenresServiceGetGenreMethodDescriptor    = moviesGenresServiceServiceDescriptor.Methods().ByName("GetGenre")
	moviesGenresServiceGetGenresMethodDescriptor   = moviesGenresServiceServiceDescriptor.Methods().ByName("GetGenres")
	moviesGenresServiceCreateGenreMethodDescriptor = moviesGenresServiceServiceDescriptor.Methods().ByName("CreateGenre")
	moviesGenresServiceUpdateGenreMethodDescriptor = moviesGenresServiceServiceDescriptor.Methods().ByName("UpdateGenre")
	moviesGenresServiceDeleteGenreMethodDescriptor = moviesGenresServiceServiceDescriptor.Methods().ByName("DeleteGenre")
)

// MoviesGenresServiceClient is a client for the movies.v1.MoviesGenresService service.
type MoviesGenresServiceClient interface {
	GetGenre(context.Context, *connect.Request[v1.GetGenreRequest]) (*connect.Response[v1.GetGenreResponse], error)
	GetGenres(context.Context, *connect.Request[v1.GetGenresRequest]) (*connect.Response[v1.GetGenresResponse], error)
	CreateGenre(context.Context, *connect.Request[v1.CreateGenreRequest]) (*connect.Response[v1.CreateGenreResponse], error)
	UpdateGenre(context.Context, *connect.Request[v1.UpdateGenreRequest]) (*connect.Response[v1.UpdateGenreResponse], error)
	DeleteGenre(context.Context, *connect.Request[v1.DeleteGenreRequest]) (*connect.Response[v1.DeleteGenreResponse], error)
}

// NewMoviesGenresServiceClient constructs a client for the movies.v1.MoviesGenresService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMoviesGenresServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MoviesGenresServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &moviesGenresServiceClient{
		getGenre: connect.NewClient[v1.GetGenreRequest, v1.GetGenreResponse](
			httpClient,
			baseURL+MoviesGenresServiceGetGenreProcedure,
			connect.WithSchema(moviesGenresServiceGetGenreMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getGenres: connect.NewClient[v1.GetGenresRequest, v1.GetGenresResponse](
			httpClient,
			baseURL+MoviesGenresServiceGetGenresProcedure,
			connect.WithSchema(moviesGenresServiceGetGenresMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createGenre: connect.NewClient[v1.CreateGenreRequest, v1.CreateGenreResponse](
			httpClient,
			baseURL+MoviesGenresServiceCreateGenreProcedure,
			connect.WithSchema(moviesGenresServiceCreateGenreMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateGenre: connect.NewClient[v1.UpdateGenreRequest, v1.UpdateGenreResponse](
			httpClient,
			baseURL+MoviesGenresServiceUpdateGenreProcedure,
			connect.WithSchema(moviesGenresServiceUpdateGenreMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteGenre: connect.NewClient[v1.DeleteGenreRequest, v1.DeleteGenreResponse](
			httpClient,
			baseURL+MoviesGenresServiceDeleteGenreProcedure,
			connect.WithSchema(moviesGenresServiceDeleteGenreMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// moviesGenresServiceClient implements MoviesGenresServiceClient.
type moviesGenresServiceClient struct {
	getGenre    *connect.Client[v1.GetGenreRequest, v1.GetGenreResponse]
	getGenres   *connect.Client[v1.GetGenresRequest, v1.GetGenresResponse]
	createGenre *connect.Client[v1.CreateGenreRequest, v1.CreateGenreResponse]
	updateGenre *connect.Client[v1.UpdateGenreRequest, v1.UpdateGenreResponse]
	deleteGenre *connect.Client[v1.DeleteGenreRequest, v1.DeleteGenreResponse]
}

// GetGenre calls movies.v1.MoviesGenresService.GetGenre.
func (c *moviesGenresServiceClient) GetGenre(ctx context.Context, req *connect.Request[v1.GetGenreRequest]) (*connect.Response[v1.GetGenreResponse], error) {
	return c.getGenre.CallUnary(ctx, req)
}

// GetGenres calls movies.v1.MoviesGenresService.GetGenres.
func (c *moviesGenresServiceClient) GetGenres(ctx context.Context, req *connect.Request[v1.GetGenresRequest]) (*connect.Response[v1.GetGenresResponse], error) {
	return c.getGenres.CallUnary(ctx, req)
}

// CreateGenre calls movies.v1.MoviesGenresService.CreateGenre.
func (c *moviesGenresServiceClient) CreateGenre(ctx context.Context, req *connect.Request[v1.CreateGenreRequest]) (*connect.Response[v1.CreateGenreResponse], error) {
	return c.createGenre.CallUnary(ctx, req)
}

// UpdateGenre calls movies.v1.MoviesGenresService.UpdateGenre.
func (c *moviesGenresServiceClient) UpdateGenre(ctx context.Context, req *connect.Request[v1.UpdateGenreRequest]) (*connect.Response[v1.UpdateGenreResponse], error) {
	return c.updateGenre.CallUnary(ctx, req)
}

// DeleteGenre calls movies.v1.MoviesGenresService.DeleteGenre.
func (c *moviesGenresServiceClient) DeleteGenre(ctx context.Context, req *connect.Request[v1.DeleteGenreRequest]) (*connect.Response[v1.DeleteGenreResponse], error) {
	return c.deleteGenre.CallUnary(ctx, req)
}

// MoviesGenresServiceHandler is an implementation of the movies.v1.MoviesGenresService service.
type MoviesGenresServiceHandler interface {
	GetGenre(context.Context, *connect.Request[v1.GetGenreRequest]) (*connect.Response[v1.GetGenreResponse], error)
	GetGenres(context.Context, *connect.Request[v1.GetGenresRequest]) (*connect.Response[v1.GetGenresResponse], error)
	CreateGenre(context.Context, *connect.Request[v1.CreateGenreRequest]) (*connect.Response[v1.CreateGenreResponse], error)
	UpdateGenre(context.Context, *connect.Request[v1.UpdateGenreRequest]) (*connect.Response[v1.UpdateGenreResponse], error)
	DeleteGenre(context.Context, *connect.Request[v1.DeleteGenreRequest]) (*connect.Response[v1.DeleteGenreResponse], error)
}

// NewMoviesGenresServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMoviesGenresServiceHandler(svc MoviesGenresServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	moviesGenresServiceGetGenreHandler := connect.NewUnaryHandler(
		MoviesGenresServiceGetGenreProcedure,
		svc.GetGenre,
		connect.WithSchema(moviesGenresServiceGetGenreMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moviesGenresServiceGetGenresHandler := connect.NewUnaryHandler(
		MoviesGenresServiceGetGenresProcedure,
		svc.GetGenres,
		connect.WithSchema(moviesGenresServiceGetGenresMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moviesGenresServiceCreateGenreHandler := connect.NewUnaryHandler(
		MoviesGenresServiceCreateGenreProcedure,
		svc.CreateGenre,
		connect.WithSchema(moviesGenresServiceCreateGenreMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moviesGenresServiceUpdateGenreHandler := connect.NewUnaryHandler(
		MoviesGenresServiceUpdateGenreProcedure,
		svc.UpdateGenre,
		connect.WithSchema(moviesGenresServiceUpdateGenreMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moviesGenresServiceDeleteGenreHandler := connect.NewUnaryHandler(
		MoviesGenresServiceDeleteGenreProcedure,
		svc.DeleteGenre,
		connect.WithSchema(moviesGenresServiceDeleteGenreMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/movies.v1.MoviesGenresService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MoviesGenresServiceGetGenreProcedure:
			moviesGenresServiceGetGenreHandler.ServeHTTP(w, r)
		case MoviesGenresServiceGetGenresProcedure:
			moviesGenresServiceGetGenresHandler.ServeHTTP(w, r)
		case MoviesGenresServiceCreateGenreProcedure:
			moviesGenresServiceCreateGenreHandler.ServeHTTP(w, r)
		case MoviesGenresServiceUpdateGenreProcedure:
			moviesGenresServiceUpdateGenreHandler.ServeHTTP(w, r)
		case MoviesGenresServiceDeleteGenreProcedure:
			moviesGenresServiceDeleteGenreHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMoviesGenresServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMoviesGenresServiceHandler struct{}

func (UnimplementedMoviesGenresServiceHandler) GetGenre(context.Context, *connect.Request[v1.GetGenreRequest]) (*connect.Response[v1.GetGenreResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("movies.v1.MoviesGenresService.GetGenre is not implemented"))
}

func (UnimplementedMoviesGenresServiceHandler) GetGenres(context.Context, *connect.Request[v1.GetGenresRequest]) (*connect.Response[v1.GetGenresResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("movies.v1.MoviesGenresService.GetGenres is not implemented"))
}

func (UnimplementedMoviesGenresServiceHandler) CreateGenre(context.Context, *connect.Request[v1.CreateGenreRequest]) (*connect.Response[v1.CreateGenreResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("movies.v1.MoviesGenresService.CreateGenre is not implemented"))
}

func (UnimplementedMoviesGenresServiceHandler) UpdateGenre(context.Context, *connect.Request[v1.UpdateGenreRequest]) (*connect.Response[v1.UpdateGenreResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("movies.v1.MoviesGenresService.UpdateGenre is not implemented"))
}

func (UnimplementedMoviesGenresServiceHandler) DeleteGenre(context.Context, *connect.Request[v1.DeleteGenreRequest]) (*connect.Response[v1.DeleteGenreResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("movies.v1.MoviesGenresService.DeleteGenre is not implemented"))
}