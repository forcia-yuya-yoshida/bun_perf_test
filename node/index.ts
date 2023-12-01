import Fastify from "fastify";
const fastify = Fastify({ logger: false });

fastify.get("/", (_, reply) => {
  reply.send("Hello world!");
});

fastify.get("/calc", (_, reply) => {
  let x = 0;
  for (let i = 0; i < 100000000; i++) {
    x++;
  }
  reply.send(x);
});

fastify.listen({ port: 3000 });
