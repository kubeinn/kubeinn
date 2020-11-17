# Build innkeeper
FROM node:15.2.0-alpine3.12 as innkeeper-builder
WORKDIR /usr/src/app
COPY ./app/innkeeper ./
RUN npm install
RUN npm run build ./build/

# Build pilgrim
FROM node:15.2.0-alpine3.12 as pilgrim-builder
WORKDIR /usr/src/app
COPY ./app/pilgrim ./
RUN npm install
RUN npm run build ./build/

# Build reeve
FROM node:15.2.0-alpine3.12 as reeve-builder
WORKDIR /usr/src/app
COPY ./app/reeve ./
RUN npm install
RUN npm run build ./build/

# Hosted on nginx
FROM nginx:1.19
RUN mkdir /usr/share/nginx/html/innkeeper
RUN mkdir /usr/share/nginx/html/pilgrim
RUN mkdir /usr/share/nginx/html/reeve
COPY --from=innkeeper-builder /usr/src/app/build /usr/share/nginx/html/innkeeper
COPY --from=pilgrim-builder /usr/src/app/build /usr/share/nginx/html/pilgrim
COPY --from=reeve-builder /usr/src/app/build /usr/share/nginx/html/reeve
EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]

# # Alternative: Integration with Shibboleth (optional)
# FROM unicon/shibboleth-sp:3.0.4
# COPY ./shibboleth-sp/ /etc/shibboleth/
# COPY ./build /var/www/html/ 