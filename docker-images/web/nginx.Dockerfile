FROM nginx:latest

# run-time variables
ENV MEDCO_END="dev"

# run
CMD /bin/bash -c "envsubst < /etc/nginx/conf.d/servers.conf.template > /etc/nginx/conf.d/servers.conf && exec nginx -g 'daemon off;'"
