package movies

import "devcodeIdentity/app/helpers/rest"

var Controller = rest.Controller{
	Prefix: "movies",
	Routes: []*rest.Route{
		{
			Method: "GET",
			Action: ActionMovies,
		},
		{
			Method: "POST",
			Action: ActionCreate,
		},
		{
			Path: "/{movie_id}",
			Method: "PUT",
			Action: ActionUpdate,
		},
		{
			Path: "/{movie_id}",
			Method: "DELETE",
			Action: ActionDelete,
		},
		{
			Path: "/{movie_id}/{actor_id}/{action}",
			Method: "POST",
			Action: ActionMovieActor,
		},
	},
}
