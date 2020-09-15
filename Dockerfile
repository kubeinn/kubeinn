# Hosted on nginx
FROM nginx:1.19
RUN mkdir /usr/share/nginx/html/innkeeper
COPY ./build /usr/share/nginx/html/innkeeper
EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]

# # Alternative: Integration with Shibboleth (optional)
# FROM unicon/shibboleth-sp:3.0.4
# COPY ./shibboleth-sp/ /etc/shibboleth/
# COPY ./build /var/www/html/ 