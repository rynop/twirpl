const M = require('./rpc/publicservices/service_pb'),
    https = require('http');

let search = new M.Search(),
    giphy = new M.Giphy();

search.setTerm("Wahooo");
const postBytes = search.serializeBinary();

const options = {
    hostname: 'localhost',
    port: 8080,
    path: '/twirp/com.rynop.twirpl.publicservices.Image/CreateGiphy',
    method: 'POST',
    headers: {
        'Content-Type': 'application/protobuf',
    }
};

const req = https.request(options, (res) => {
    console.log('statusCode:', res.statusCode);
    console.log('headers:', res.headers);

    res.on('data', (d) => {
        const uInt8Array = d.buffer.slice(d.byteOffset, d.byteOffset + d.byteLength);
        const giphyMessage = M.Giphy.deserializeBinary(uInt8Array);
        console.log(giphyMessage.toObject());
    });
});

req.on('error', (e) => {
    console.error(e);
});

req.write(Buffer.from(postBytes));
req.end();
