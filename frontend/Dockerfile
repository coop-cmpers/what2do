FROM node:20-alpine

RUN mkdir /app
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY yarn.lock ./

#RUN rm -rf /app/node_modules
RUN yarn

COPY . ./

EXPOSE 3000

CMD ["yarn", "start"]