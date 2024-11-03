gcloud run deploy vicyberapi \
  --source . \
  --region europe-west3 \
  --allow-unauthenticated \
  --set-env-vars VICYBERAPIKEY=vicyberthebest,POSTGRES_URL=postgres://default:bPR6ofNAKJ5C@ep-super-flower-a2vcv1zv-pooler.eu-central-1.aws.neon.tech:5432/verceldb?sslmode=require
