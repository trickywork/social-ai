# SocialAI

SocialAI is a Go backend with a compact web demo for a social media search/upload project. The portfolio deployment runs in demo mode so it can be hosted cheaply on Cloud Run without Elasticsearch or Google Cloud Storage.

## Live Demo

- Portfolio URL: `https://socialai.junliu.dev`
- Cloud Run service: `socialai`
- Cloud Run URL: `https://socialai-gb7rmueyna-uc.a.run.app`
- Google Cloud project: `caramel-vim-441513-e1`
- Region: `us-central1`

The custom domain mapping exists in Cloud Run. If the domain is still pending in Google Cloud Console, verify the Cloudflare DNS record and wait for the Google-managed certificate to finish provisioning.

## Tech Stack

- Go
- Gorilla Mux
- JWT authentication
- Multipart file upload
- In-memory demo store
- Elasticsearch adapter for full mode
- Google Cloud Storage adapter for full mode
- Docker, Google Cloud Build, Google Cloud Run
- API testing via local Postman workspace

## Project Structure

```text
SocialAI/
  socialai/
    handler/
    service/
    model/
    conf/
    web/
    Dockerfile
  docs/
    configuration.md
  cloudbuild.yaml
```

## Features

- Signup and signin.
- JWT-protected routes.
- Upload a post with text and media.
- Search posts by username or keyword.
- Delete a post owned by the current user.
- Static web UI served by the Go server.
- Demo mode that avoids external infrastructure.

## Local Development

Run in demo mode:

```bash
cd /Users/junliu/git_repo/SocialAI/socialai
SOCIALAI_MODE=demo TOKEN_SECRET=dev-secret go run .
```

Expected local URL:

```text
http://localhost:8080
```

Health check:

```bash
curl http://localhost:8080/api/health
```

Expected result:

```json
{"status":"ok"}
```

## Environment Variables

```env
PORT=8080
SOCIALAI_MODE=demo
TOKEN_SECRET=dev-secret
```

Full mode also needs external service configuration in:

```text
socialai/conf/deploy.yml
```

The checked-in config uses placeholders. Do not commit real GCS, Elasticsearch, or credential values.

## API Endpoints

| Method | Path | Auth | Description |
| --- | --- | --- | --- |
| `GET` | `/api/health` | No | Health check. |
| `POST` | `/signup` | No | Register a user. |
| `POST` | `/signin` | No | Sign in and return a JWT token. |
| `POST` | `/upload` | Yes | Upload media and create a post. |
| `GET` | `/search?user={username}` | Yes | Search posts by username. |
| `GET` | `/search?keywords={text}` | Yes | Search posts by keywords. |
| `DELETE` | `/post/{id}` | Yes | Delete a post owned by the signed-in user. |

## Postman

Use the local Postman workspace collections:

```text
SocialAI Backend 01. Go Introduction
SocialAI Backend 02. Elasticsearch
SocialAI Backend 03. GCS
202409 SocialAI - Coding Pad API Tests
```

Suggested variables:

```text
baseUrl=http://localhost:8080
token=
```

The collection stores the signin token in a collection variable after a successful signin request.

For Cloud Run:

```text
baseUrl=https://socialai-gb7rmueyna-uc.a.run.app
```

The repo-exported backup copy is stored outside GitHub at:

```text
/Users/junliu/CourseArtifacts/postman/project-exported/SocialAI.postman_collection.json
```

## Tests And Build

```bash
cd /Users/junliu/git_repo/SocialAI/socialai
go test ./...
go build ./...
```

## Cloud Deployment

Manual deployment:

```bash
cd /Users/junliu/git_repo/SocialAI
gcloud builds submit \
  --config cloudbuild.yaml \
  --project caramel-vim-441513-e1
```

Cloud Run cost controls:

- `min-instances=0`
- `max-instances=1`
- `SOCIALAI_MODE=demo`
- no Elasticsearch VM
- no GCS writes in demo mode

Runtime settings:

```text
SOCIALAI_MODE=demo
TOKEN_SECRET=Secret Manager: socialai-token-secret
```

## Expected Portfolio Behavior

A visitor should be able to open the demo UI, create an account, sign in, create a sample post, search posts, and delete their own post. In demo mode, data is temporary and resets when the Cloud Run instance restarts.

## Additional Notes

Configuration details are in:

```text
docs/configuration.md
```
