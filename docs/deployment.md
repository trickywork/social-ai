# SocialAI Deployment

Cloud Run service:

```text
socialai
```

Current URL:

```text
https://socialai-gb7rmueyna-uc.a.run.app
```

GitHub trigger:

```text
socialai-main-deploy
```

## Manual Deploy

```bash
gcloud builds submit --config cloudbuild.yaml --project caramel-vim-441513-e1
```

## Secret Manager

Cloud Run receives `TOKEN_SECRET` from:

```text
socialai-token-secret
```

## Demo Mode

Demo mode is set with:

```text
SOCIALAI_MODE=demo
```

This replaces Elasticsearch and GCS with in-memory storage for a cheap portfolio deployment. Data resets when the Cloud Run instance restarts.

## Full Mode

For a production-like version, configure:

```text
socialai/conf/deploy.yml
```

Then deploy without `SOCIALAI_MODE=demo`. This may require paid resources such as Elasticsearch and Google Cloud Storage.
