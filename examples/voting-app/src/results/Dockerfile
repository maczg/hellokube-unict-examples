FROM python:3.10-buster

WORKDIR /app

COPY ./app .
COPY ./requirements.txt .

RUN pip install -r requirements.txt

CMD ["flask", "run"]