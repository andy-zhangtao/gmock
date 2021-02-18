FROM vikings/alpine
COPY bin/gmock /gmock
ENTRYPOINT ["/gmock"]