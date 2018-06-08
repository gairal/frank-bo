install:
	dep ensure

run:
	go run main.go

deploy:
	gcloud app deploy --project com-gairal-frank-api
