FROM node:19.9.0-bullseye


WORKDIR /usr/src/app

#Copiamos el codigo de nuestra app
COPY . .

#Instalamos las dependencias
RUN npm install -g npm
RUN npm install

#Ejecutamos nuestra app en docker
ENTRYPOINT [ "npm", "start" ]