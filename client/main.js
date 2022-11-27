const prompt = require("prompt-sync")({ sigint: true })
const proxy = require("./proxy")
const yargs = require('yargs/yargs')
const { hideBin } = require('yargs/helpers')


// Flags
const argv = yargs(hideBin(process.argv))
  .option('proxy-url', {
    alias: 'u',
    type: 'string',
    description: 'Proxy url (address:port)'
  })
  .option('parallel', {
    alias: 'p',
    type: 'number',
    default: 10,
    description: 'Number of maximum parallel requests. Default 10'
  })
  .option('target', {
    alias: 't',
    type: 'string',
    description: 'Target url which can be set multiple times'
  })
  .help('h')
  .alias('h', "help")
  .argv

for (let i = 0; i < argv.parallel; i++) {
    proxy.callRandomly(argv.proxyUrl, argv.target)
}
