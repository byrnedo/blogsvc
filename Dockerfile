FROM scratch

COPY main /

COPY conf /conf

EXPOSE 80

ENTRYPOINT ["/main"]


