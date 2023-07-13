# syntax = docker/dockerfile:latest

FROM alpine AS base
COPY bot /bot
EXPOSE 50051

ENTRYPOINT [ "/bot" ]
