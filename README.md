# Huma API Example

I checked out Huma because I wanted...needed to get OpenAPI documentation for a proof of concept.  Additionally,
I needed to handle api versioning and versioning of the underlying models.  So this is a proof of concept for that. I've
tried to keep dependencies few and only ones that I have a good reason to use. 


## Goals
### OpenAPI Documentation
Needed for the proof of concept, like a business requirement ;) 

### Versioned API and Models
Versioning an API is just good practice, but there were issues trying to do this with Huma

## Dependencies
### Huma
I used Huma when I saw it was easy to make an OpenAPI doc generated programmatically from the API itself, similarly to 
how things like Swashbuckle in .NET works.  This was a piece that has been missing for many of my Go-based APIs and 
when I want to advocate for Go in a more "enterprise" environment, this will help. 

### Gorilla Mux
I love this package. I was so happy when they pulled it out of archive mode and started working on it again. I briefly 
considered trying to work on it. I would like to try to make this use the stdlib's http.ServeMux library. However, 
Gorilla Mux's subrouter's allows this to work. 

### Testify
Another package I love because it simplifies the testing process.  I like to use it for mocking and assertions, makes 
them simple and clean. It has low dependency overhead and is widely used.

## Organization
- `/api` - The handlers and routes for each version of the api
- `/models` - The models that would be used by the service to handle data
- `/service` - Service focused bits, such as the api versioning and loading components
- `/cmd/service` - Cmd to run the service, I like to split this out in case I want to use bits for a cli later

## Comments?
Just an experiment, happy to take PRs if you think there is a better way. I'm happy to see what else can be done!

### To Do
- [ ] Generate OpenAPI Docs for validation, ie Pact.io
- [ ] Tests
- [ ] GitHub Actions to run validation and tests