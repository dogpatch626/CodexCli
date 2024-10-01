

const fastify = require("fastify")
const cors = require("@fastify/cors")
const fastifyStatic = require("@fastify/static")
const path = require('path')

let server = fastify({ logger: true });
server.register(cors)
server.register(fastifyStatic, { root: path.join(__dirname, 'static'), prefix: "/static/" })

server.get('/', (request, reply) => reply.status(200).sendFile('index.html'))

server.listen({ port: 3000 }, function (err, address) {
    if (err) {
        server.log.error(err)
        console.log(err)
    }
});
