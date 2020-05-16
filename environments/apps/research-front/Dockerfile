FROM google/dart AS build-env

RUN pub global activate webdev

COPY ./sources/ /work/
WORKDIR /work

RUN pub get
RUN webdev build -- --delete-conflicting-outputs

FROM nginx:mainline-alpine

COPY --from=build-env \
  /work/build/favicon.* \
  /work/build/*.html \
  /work/build/*.js \
  /work/build/*.css \
  /usr/share/nginx/html/
COPY --from=build-env \
 /work/build/images/ \
 /usr/share/nginx/html/images/

COPY nginx.conf /etc/nginx/
COPY config-template.conf /etc/nginx/conf.d/

COPY run.sh /usr/local/bin/run.sh
RUN chmod +x /usr/local/bin/run.sh

ENTRYPOINT /usr/local/bin/run.sh
