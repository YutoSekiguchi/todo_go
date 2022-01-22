FROM --platform=linux/x86_64 node:15.8


WORKDIR /app

COPY . .
RUN npm install


CMD [ "npm", "run", "serve" ]