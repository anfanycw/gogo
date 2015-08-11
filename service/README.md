# GOGO Service Setup

1. Run `go get github.com/gorilla/mux` inside gogo folder. Make sure `$GOPATH` and `$GOBIN` are set and exists in `$PATH`
2. Run `go build && ./service` to start the gogo service


# Endpoints

## POST /occupancy

    Post Body
    { “status”: ”VACANT” | ”OCCUPIED” }


## GET /occupancy

    Response
    {
        “vacant”: [  “upstairs left”, “upstairs right”]
    }


## POST /reservation

    Request Body
    {
        “name”: “who am i”,
        “location”: “upstairs left”
    }

* Response 404 if location is not found
* Response 400 if location is already occupied
* Response 200 if reservation is successful


## GET /reservation/{location}

    Response
    {
        “status”: ”RESERVED”
        “name”: “my name”
    }

    {
        “status”: ”VACANT” | ”OCCUPIED”
    }
