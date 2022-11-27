var http = require('http');
module.exports = { callRandomly };

function httpGet(url) {
    return new Promise((resolve, reject) => {
        http.get(url, res => {
        res.setEncoding('utf8');
        let body = ''; 
        res.on('data', chunk => body += chunk);
        res.on('end', () => resolve([res.statusCode, body.length]));
        }).on('error', reject);
    });
}

async function callRandomly(proxyURL, targetURLs) 
{
    while(1)
    {
        try
        {
            const targetURL = targetURLs[Math.floor(Math.random()*targetURLs.length)]
            const URL = "http://" + proxyURL + "/" + targetURL
            var start = new Date();
            const res = await httpGet(URL);
            console.log("---------")
            console.log("called url:", targetURL)
            console.log("status:", res[0])
            console.log("length:", res[1])
            console.log("duration(ms):", new Date() - start)
            console.log("---------\n")
        }
        catch (exceptionVar)
        {
            console.log("exception:", exceptionVar)
        }
        await sleep(3000)
    }
}

function sleep(ms) {
    return new Promise((resolve) => {
        setTimeout(resolve, ms);
    });
}

