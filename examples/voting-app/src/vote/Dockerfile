FROM python:3.11-alpine

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

ENV FLASK_RUN_HOST="0.0.0.0"
ENV FLASK_RUN_PORT="8080"

ENTRYPOINT [ "flask", "run" ]