FROM ubuntu
COPY ./api ~/
ENV REGION=us-west-1
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates
CMD ["~/api"]