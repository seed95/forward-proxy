const prompt = require("prompt-sync")({ sigint: true })
const proxy = require("./proxy")

// const proxyURL = prompt("Enter proxy URL: ")
const proxyURL = "localhost:2121"

// let targetURLs = prompt("Enter target URLs separated by commas: ")
let targetURLs = "https://www.google.com,https://www.digitalocean.com/community/tutorials/how-to-create-an-http-client-with-core-http-in-node-js"
targetURLs = targetURLs.split(",")

const maxParallelRequests = prompt("Enter the number of maximum parallel requests: ")
// const maxParallelRequests = 2


for (let i = 0; i < maxParallelRequests; i++) {
    proxy.callRandomly(proxyURL, targetURLs)
}



// const myArgs = process.argv.slice(2);
// console.log('myArgs:', myArgs[0], myArgs[1]);


// const commander = require('commander');

// commander
//   .version('1.0.0', '-v, --version')
//   .usage('[OPTIONS]...')
//   .option('-u, --proxy-url <value>', 'Proxy url(address:port).')
//   .option('-c, --custom <value>', 'Overwriting value.', 'Default')
//   .parse(process.argv);

// const options = commander.opts();

// const flag = (options.flag ? 'Flag is present.' : 'Flag is not present.');

// console.log('Flag:', `${flag}`);
// console.log('Custom:', `${options.custom}`);