# Build crossroads
FROM node:15.2.0-alpine3.12 as crossroads-builder
ENV NODE_ENV=production
WORKDIR /usr/src/app
COPY ["./client/crossroads/package.json", "./client/crossroads/package-lock.json*", "./"]
RUN npm install
COPY ./client/crossroads ./
RUN npm run build ./build/

# Build innkeeper
FROM node:15.2.0-alpine3.12 as innkeeper-builder
ENV NODE_ENV=production
WORKDIR /usr/src/app
COPY ["./client/innkeeper/package.json", "./client/innkeeper/package-lock.json*", "./"]
RUN npm install
COPY ./client/innkeeper ./
RUN npm run build ./build/

# Build pilgrim
FROM node:15.2.0-alpine3.12 as pilgrim-builder
ENV NODE_ENV=production
WORKDIR /usr/src/app
COPY ["./client/pilgrim/package.json", "./client/pilgrim/package-lock.json*", "./"]
RUN npm install
COPY ./client/pilgrim ./
RUN npm run build ./build/

# Build schutterij backend
FROM golang:alpine3.12
RUN mkdir /app
RUN mkdir -p /app/client/innkeeper/build
RUN mkdir -p /app/client/pilgrim/build
RUN mkdir -p /app/client/crossroads/build
COPY --from=crossroads-builder /usr/src/app/build /app/client/crossroads/build
COPY --from=innkeeper-builder /usr/src/app/build /app/client/innkeeper/build
COPY --from=pilgrim-builder /usr/src/app/build /app/client/pilgrim/build
RUN mkdir /app-builder
ADD . /app-builder/
WORKDIR /app-builder
RUN go build -o . ./cmd/main.go
RUN mv ./main ../app/main
WORKDIR /app
RUN rm -r /app-builder
CMD [ "/app/main" ]

# # Hosted on nginx
# FROM nginx:1.19
# RUN mkdir /root/schutterij
# RUN mkdir /root/schutterij/client/innkeeper/build
# RUN mkdir /root/schutterij/client/pilgrim/build
# RUN mkdir /root/schutterij/client/crossroads/build
# COPY --from=crossroads-builder /usr/src/app/build /root/schutterij/client/crossroads/build
# COPY --from=innkeeper-builder /usr/src/app/build /root/schutterij/client/innkeeper/build
# COPY --from=pilgrim-builder /usr/src/app/build /root/schutterij/client/pilgrim/build
# COPY --from=schutterij-builder /app/main /home/schutterij/main
# EXPOSE 80
# ENTRYPOINT ["nginx", "-g", "daemon off;"]

# # Alternative: Integration with Shibboleth (optional)
# FROM unicon/shibboleth-sp:3.0.4
# COPY ./shibboleth-sp/ /etc/shibboleth/
# COPY ./build /var/www/html/ 