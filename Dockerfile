FROM scratch

EXPOSE 8080

COPY k8s-demo /

CMD ["./k8s-demo", "-port", "8080"]