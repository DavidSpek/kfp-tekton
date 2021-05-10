FROM python:3.9.5-slim

RUN pip install kubernetes Flask flask-cors requests

ENV APP_HOME /app
COPY src $APP_HOME
WORKDIR $APP_HOME

ENTRYPOINT ["python", "serve.py"]
