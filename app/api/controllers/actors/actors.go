package actors

import "devcodeIdentity/app/helpers/rest"

var Controller = rest.Controller{
	Prefix: "actors",
	Routes: []*rest.Route{
		{
			Method: "GET",
			Action: ActionActors,
		},
		{
			Method: "POST",
			Action: ActionCreate,
		},
		{
			Path: "/{actor_id}",
			Method: "PUT",
			Action: ActionUpdate,
		},
		{
			Path: "/{actor_id}",
			Method: "DELETE",
			Action: ActionDelete,
		},
	},
}
