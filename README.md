# user-segmentation

## Build
#### Prerequisites
- docker

1. Clone project:
```
git clone https://github.com/vlasashk/user-segmentation.git
cd user-segmentation
```
2. Run:
```
docker compose up --build
```
## Project information
API for dynamic user segmentation for testing new functionality
### Restrictions
- UserID must be unique, otherwise it won't be allowed to add one
- Segment name must be unique, otherwise it won't be allowed to add one
- User can be added to segment only if user and segment both exist in database
- All segments present in the list to be added to a user must exist in the database.
Even if a single segment doesn't exist then the request will be aborted and none of the segments will be added to a user
- If segment is already assigned to user then the request will be aborted and none of the segments from the list will be added to a user
- Deleting segment from database will cascade delete it from every user and history for this segment won't be available
- Deleting segment from a user doesn't delete record from database, instead of deletion it marks `deleted_at` field with current date
- In order to remove user from segment - all segments from the request must be present in database
- If user is not part of some segments from the removal list then the removal request will be aborted and none of the segments from the list will be removed from a user

### Tools used
- PostgreSQL as database
- [jackc/pgx](https://pkg.go.dev/github.com/jackc/pgx) package as toolkit for PostgreSQL
- [go-chi/chi](https://pkg.go.dev/github.com/go-chi/chi) package as router for building HTTP service
- [swaggo/swag](https://github.com/swaggo/swag) package as swagger doc generator
- Docker for deployment

### Functionality
#### Swagger
Swagger generated documentation will be available after run at `http://localhost:8090/swagger/index.html` (or different port if .env file was edited)

#### Users manipulation
- {POST} **/user/new** - Add new user to database.</br> Request Body JSON:
```
{
    "user_id": 10
}
```
- {POST} **/user/addSegment** - Add list of segments to user. 
User and each segment must be present in database for successful execution 
otherwise it won't ve allowed.</br> Request Body JSON:
```
{
    "user_id": 10,
    "segment_slug": ["AVITO", "AVITO_10", "AVITO_30"]
}
```
- {GET} **/user/segments/{userID}** - Return the list of segments the user is a member of.</br> Request Body is not required.
- {DELETE} **/user/segments** - Remove user from chosen segments by marking deleted_at field.</br> Request Body JSON:
```
{
    "user_id": 10,
    "segment_slug": ["AVITO", "AVITO_10", "AVITO_30"]
}
```

#### Segments manipulation
- {POST} **/segment/new** Add new segment to database.</br> Request Body JSON:
```
{
    "slug": "test"
}
```
- {DELETE} **/segment/remove**  Cascade delete segment. 
This method will permanently delete segment and all it's relations between user-segment.</br> Request Body JSON:
```
{
    "slug": "test"
}
```
- {GET} **/segment/users/{segmentName}** - Return the list of users the segment has.</br> Request Body is not required.
#### CSV Report
- {POST} **/report** - Return the link to csv file with report for chosen month.</br> Request Body JSON:
```
{
    "year": 2023,
    "month": 9
}
```
- {GET} **/report/{fileName}** - Download the csv file with report for chosen month.</br> Request Body is not required.