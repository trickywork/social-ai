# SocialAI Configuration

This file records the non-code setup needed to run, test, and redeploy SocialAI.

## Runtime Shape

SocialAI is a Go backend that also serves a compact static portfolio UI.

The project has two storage modes:

- `demo` / `memory`: in-memory users and posts, no Elasticsearch, no Google Cloud Storage.
- full mode: Elasticsearch for post/user search plus Google Cloud Storage for media.

The current Cloud Run deployment uses demo mode to keep cost low. Data resets when the Cloud Run instance restarts.

## Local Environment

Use `.env.example` as the template:

```env
PORT=8080
SOCIALAI_MODE=demo
TOKEN_SECRET=dev-secret
```

Local run:

```bash
cd /Users/junliu/git_repo/SocialAI/socialai
SOCIALAI_MODE=demo TOKEN_SECRET=dev-secret go run .
```

Health check:

```bash
curl http://localhost:8080/api/health
```

## Demo Storage

Demo mode initializes:

```text
socialai/backend/memory.go
```

No local database is required. The in-memory maps store:

- users
- posts

Uploads are represented by metadata for portfolio testing; the demo avoids paid GCS writes.

## Full Course Storage

Full mode reads:

```text
socialai/conf/deploy.yml
```

Template fields:

```yaml
elasticsearch:
    address: "http://YOUR_ELASTICSEARCH_HOST:9200/"
    username: "YOUR_ELASTICSEARCH_USERNAME"
    password: "YOUR_ELASTICSEARCH_PASSWORD"

gcs:
    bucket: "YOUR_GCS_BUCKET"

token:
    secret: "change-me"
```

Do not commit real Elasticsearch credentials, GCS credentials, or JWT secrets.

## API Testing

Postman collection:

```text
postman/SocialAI.postman_collection.json
```

Variables:

```text
baseUrl=http://localhost:8080
token=<auto-filled after signin>
```

Test order:

1. `POST /signup`
2. `POST /signin`
3. `POST /upload`
4. `GET /search`
5. `DELETE /post/{id}`

## Cloud Resources

Google Cloud project:

```text
caramel-vim-441513-e1
```

Region:

```text
us-central1
```

Cloud Run service:

```text
socialai
```

Cloud Run URL:

```text
https://socialai-gb7rmueyna-uc.a.run.app
```

Custom domain mapping:

```text
socialai.junliu.dev
```

Cloud Build trigger:

```text
socialai-main-deploy
```

Secret Manager:

```text
socialai-token-secret -> TOKEN_SECRET
```

Cloud Run env vars:

```text
SOCIALAI_MODE=demo
```

## Cost Notes

- Current demo mode has no Elasticsearch VM, no Cloud SQL, and no GCS write path.
- Cloud Run is configured for `min-instances=0`.
- Full mode will cost more because Elasticsearch and object storage need persistent resources.
