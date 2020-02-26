# python version to change
FROM python:3.8-alpine


# set work directory

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN apk update \
    && apk add --virtual build-deps gcc python3-dev musl-dev \
    && apk add postgresql-dev \
    && pip install psycopg2 \
    && apk del build-deps

# install dependencies
RUN pip install --upgrade pip
RUN pip install -r /usr/src/app/matcha/requirements.txt

# copy project
COPY matcha /usr/src/app/

# run entrypoint.sh
ENTRYPOINT ["/usr/src/app/matcha/entrypoint.sh"]
