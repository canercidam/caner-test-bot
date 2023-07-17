# syntax = docker/dockerfile:latest

FROM alpine AS base
COPY bot /bot
EXPOSE 50051

LABEL "network.forta.settings.agent-logs.enable"="true"

ENTRYPOINT [ "/bot" ]
