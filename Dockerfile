FROM scratch
COPY api-scenario /
ENTRYPOINT ["/api-scenario"]