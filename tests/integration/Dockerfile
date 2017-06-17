FROM node:7

RUN mkdir -p /workspace/tests/integration
WORKDIR /workspace/tests/integration

ADD ./package.json ./package.json
RUN npm install

ENTRYPOINT []
CMD ["npm", "test"]
