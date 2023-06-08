FROM node:alpine

WORKDIR /app

ENV PATH /view/node_modules/.bin:$PATH

COPY ./view/package*.json /app/.

RUN npm install

COPY ./view/* .

EXPOSE 3000

CMD ["npm" "run" "start" ]