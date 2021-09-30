#A simple rest api server using go and gin that provides access to a store selling vintage vinyl records.
###API spec:-
>/albums
>
> - GET - get a list of all albums,returned as JSON
>
> - POST - add a new album from request data sent as JSON
>
>/albums/:id
>
> - GET - Get an album data by ID ,returns album data as JSON
