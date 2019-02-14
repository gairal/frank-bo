install:
	dep ensure

run:
	go run main.go

dev:
	dev_appserver.py appengine/app.yaml

deploy:
	gcloud app deploy ./appengine/app.yaml --project com-gairal-frank
