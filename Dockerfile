FROM ubuntu

WORKDIR /tune-bot/

RUN apt-get update && apt-get install -y dos2unix git

COPY . ./api
RUN git clone https://github.com/tune-bot/core.git
RUN dos2unix api/infrastructure/install.sh
RUN dos2unix core/infrastructure/install.sh
RUN dos2unix core/infrastructure/create.sql
RUN mv api/infrastructure/database.env core/infrastructure/database.env

RUN chmod +x core/infrastructure/install.sh

RUN core/infrastructure/install.sh
RUN api/infrastructure/install.sh

RUN rm -rf core && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN bin/database
RUN bin/api
EXPOSE 80

ENTRYPOINT ["tail", "-f", "/dev/null"]
