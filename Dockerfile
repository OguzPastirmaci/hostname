FROM node:7-alpine

WORKDIR /src
ADD . .

RUN npm install

EXPOSE 8000

CMD ["node", "index.js"] 
