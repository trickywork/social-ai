# SocialAI

Go backend and compact web demo for the LaiOffer Social AI project. The portfolio deployment keeps the course API shape while running in low-cost demo mode on Cloud Run.

## Live Service

- Cloud Run service: `socialai`
- Current URL: `https://socialai-gb7rmueyna-uc.a.run.app`
- GitHub trigger: `socialai-main-deploy`

## Features

- Signup and signin with JWT.
- Upload a post with media.
- Search posts by user or keywords.
- Delete your own post.
- Demo mode uses an in-memory backend so it does not require Elasticsearch or Google Cloud Storage.
- Static web UI is served by the Go server for portfolio viewing.

## Tech Stack

- Go
- Gorilla Mux
- JWT middleware
- Elasticsearch and GCS adapters from the course project
- In-memory demo backend for Cloud Run

## Local Run

```bash
cd socialai
SOCIALAI_MODE=demo TOKEN_SECRET=dev-secret go run .
```

Open:

```text
http://localhost:8080
```

Health check:

```bash
curl http://localhost:8080/api/health
```

## Tests

```bash
cd socialai
go test ./...
go build ./...
```

## API Endpoints

| Method | Path | Auth | Description |
| --- | --- | --- | --- |
| GET | `/api/health` | No | Health check. |
| POST | `/signup` | No | Register a user. |
| POST | `/signin` | No | Sign in and return a JWT token. |
| POST | `/upload` | Yes | Upload media and create a post. |
| GET | `/search?user={username}` | Yes | Search posts by user. |
| GET | `/search?keywords={text}` | Yes | Search posts by keywords. |
| DELETE | `/post/{id}` | Yes | Delete one post owned by the signed-in user. |

## API Testing

Import `postman/SocialAI.postman_collection.json`.

The collection stores the signin token in a `token` collection variable.

## Configuration Notes

Non-code setup is documented in `docs/configuration.md`, including demo memory mode, full Elasticsearch/GCS mode, Secret Manager, and Cloud Run settings.

## Cloud Run Deployment

The repo includes `socialai/Dockerfile` and root `cloudbuild.yaml`.

Manual deploy:

```bash
gcloud builds submit --config cloudbuild.yaml --project caramel-vim-441513-e1
```

Runtime settings:

```text
SOCIALAI_MODE=demo
TOKEN_SECRET=Secret Manager: socialai-token-secret
```

Cost controls:

- Cloud Run `min-instances=0`.
- Cloud Run `max-instances=1`.
- No Elasticsearch VM in demo mode.
- No GCS bucket writes in demo mode.

## Full Course Mode

To run with Elasticsearch and GCS, set `SOCIALAI_MODE` to a non-demo value and configure `socialai/conf/deploy.yml`. The checked-in `deploy.yml` intentionally uses placeholders instead of real credentials.
